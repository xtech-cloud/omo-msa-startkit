package config

import (
	"encoding/json"
	"os"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/yaml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

var defaultConfigDefine string = `
	{	
		"source": "file",
		"prefix": "./_runpath/",
		"key": "default.yaml"
	}	
`

type ConfigDefine struct {
	Source  string   `json:source`
	Prefix  string   `json:prefix`
	Key     string   `json:key`
	Address []string `json:address`
}

var Scheme ConfigScheme_

func Setup() {
	conf, err := config.NewConfig()
	if nil != err {
		panic(err)
	}

	//registry plugin
	registryPlugin := os.Getenv("MSA_REGISTRY_PLUGIN")
	if "" == registryPlugin {
		registryPlugin = "consul"
	}
	os.Setenv("MICRO_REGISTRY", registryPlugin)

	//registry address
	registryAddress := os.Getenv("MSA_REGISTRY_ADDRESS")
	if "" == registryAddress {
		registryPlugin = "127.0.0.1:8500"
	}
	os.Setenv("MICRO_REGISTRY_ADDRESS", registryAddress)

	//config
	envConfigDefine := os.Getenv("MSA_CONFIG_DEFINE")
	if "" == envConfigDefine {
		envConfigDefine = defaultConfigDefine
	}
	logger.Infof("config define as %v", envConfigDefine)

	var configDefine ConfigDefine
	err = json.Unmarshal([]byte(envConfigDefine), &configDefine)
	if err != nil {
		panic(err)
		return
	}

	if "file" == configDefine.Source {
		filepath := configDefine.Prefix + configDefine.Key
		fileSource := file.NewSource(
			file.WithPath(filepath),
		)
		err = conf.Load(fileSource)
		if nil != err {
			logger.Errorf("load config %v failed: %v", filepath, err)
			panic(err)
		}
		logger.Infof("load config %v success", filepath)
		conf.Scan(&Scheme)
	} else if "consul" == configDefine.Source {
		consulKey := configDefine.Prefix + configDefine.Key
		for _, addr := range configDefine.Address {
			consulSource := consul.NewSource(
				consul.WithAddress(addr),
				consul.WithPrefix(configDefine.Prefix),
				consul.StripPrefix(true),
				source.WithEncoder(yaml.NewEncoder()),
			)
			err = conf.Load(consulSource)
			if nil == err {
				logger.Infof("load config %v from %v success", consulKey, addr)
				break
			} else {
				logger.Errorf("load config %v from %v failed: %v", consulKey, addr, err)
			}
		}
		conf.Get(configDefine.Key).Scan(&Scheme)
	} else if "etcd" == configDefine.Source {
		consulKey := configDefine.Prefix + configDefine.Key
		for _, addr := range configDefine.Address {
			etcdSource := etcd.NewSource(
				etcd.WithAddress(addr),
				etcd.WithPrefix(configDefine.Prefix),
				etcd.StripPrefix(true),
				source.WithEncoder(yaml.NewEncoder()),
			)
			err = conf.Load(etcdSource)
			if nil == err {
				logger.Infof("load config %v from %v success", consulKey, addr)
			} else {
				logger.Errorf("load config %v from %v failed: %v", consulKey, addr, err)
			}
		}
		conf.Get(configDefine.Key).Scan(&Scheme)
	}

	logger.Infof("logger level is %v", Scheme.Logger.Level)
}
