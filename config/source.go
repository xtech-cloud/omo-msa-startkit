package config

import (
	"os"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/yaml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

type RegistrySection struct {
	Plugin  string `yaml:plugin`
	Address string `yaml:address`
}

type ConfigScheme struct {
	Registry RegistrySection `yaml:registry`
}

var Scheme ConfigScheme

func Setup() {
	conf, err := config.NewConfig()
	if nil != err {
		panic(err)
	}

	configSource := os.Getenv("STARTKIT_CONFIG_SOURCE")
	configPath := os.Getenv("STARTKIT_CONFIG_PATH")

	if "consul" == configSource {
		if "" == configPath {
			configPath = "default"
		}
		address := os.Getenv("STARTKIT_CONFIG_CONSUL_ADDRESS")
		if "" == address {
			address = "127.0.0.1:8500"
		}
		prefix := os.Getenv("STARTKIT_CONFIG_CONSUL_PREFIX")
		if "" == prefix {
			prefix = "/msa/startkit/config"
		}
		consulSource := consul.NewSource(
			consul.WithAddress(address),
			consul.WithPrefix(prefix),
			consul.StripPrefix(true),
			source.WithEncoder(yaml.NewEncoder()),
		)
		loadErr := conf.Load(consulSource)
		if nil != loadErr {
			panic(loadErr)
		}
		conf.Get(configPath).Scan(&Scheme)
	} else {
		if "" == configPath {
			configPath = "./_runpath/default.yaml"
		}
		fileSource := file.NewSource(
			file.WithPath(configPath),
		)
		loadErr := conf.Load(fileSource)
		if nil != loadErr {
			panic(loadErr)
		}
		conf.Scan(&Scheme)
	}

	os.Setenv("MICRO_REGISTRY", Scheme.Registry.Plugin)
	os.Setenv("MICRO_REGISTRY_ADDRESS", Scheme.Registry.Address)
}
