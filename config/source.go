package config

import (
	"encoding/json"
	"os"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/encoder/yaml"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/config/source/memory"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"github.com/micro/go-plugins/logger/zerolog/v2"
	goYAML "gopkg.in/yaml.v2"
)

type ConfigDefine struct {
	Source  string   `json:source`
	Prefix  string   `json:prefix`
	Key     string   `json:key`
	Address []string `json:address`
}

var configDefine ConfigDefine

var Schema ConfigSchema_

func setupEnvironment() {
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
		logger.Warn("MSA_CONFIG_DEFINE is empty")
		return
	}

	logger.Infof("MSA_CONFIG_DEFINE is %v", envConfigDefine)
	err := json.Unmarshal([]byte(envConfigDefine), &configDefine)
	if err != nil {
		logger.Error(err)
	}
}

func mergeFile(_config config.Config) {
	filepath := configDefine.Prefix + configDefine.Key
	fileSource := file.NewSource(
		file.WithPath(filepath),
	)
	err := _config.Load(fileSource)
	if nil != err {
		logger.Errorf("load config %v failed: %v", filepath, err)
	} else {
		logger.Infof("load config %v success", filepath)
		_config.Scan(&Schema)
	}
}

func mergeConsul(_config config.Config) {
	consulKey := configDefine.Prefix + configDefine.Key
	for _, addr := range configDefine.Address {
		consulSource := consul.NewSource(
			consul.WithAddress(addr),
			consul.WithPrefix(configDefine.Prefix),
			consul.StripPrefix(true),
			source.WithEncoder(yaml.NewEncoder()),
		)
		err := _config.Load(consulSource)
		if nil == err {
			logger.Infof("load config %v from %v success", consulKey, addr)
			break
		} else {
			logger.Errorf("load config %v from %v failed: %v", consulKey, addr, err)
		}
	}
	_config.Get(configDefine.Key).Scan(&Schema)
}

func mergeEtcd(_config config.Config) {
	etcdKey := configDefine.Prefix + configDefine.Key
	for _, addr := range configDefine.Address {
		etcdSource := etcd.NewSource(
			etcd.WithAddress(addr),
			etcd.WithPrefix(configDefine.Prefix),
			etcd.StripPrefix(true),
			source.WithEncoder(yaml.NewEncoder()),
		)
		err := _config.Load(etcdSource)
		if nil == err {
			logger.Infof("load config %v from %v success", etcdKey, addr)
			break
		} else {
			logger.Errorf("load config %v from %v failed: %v", etcdKey, addr, err)
		}
	}
	_config.Get(configDefine.Key).Scan(&Schema)
}

func Setup() {
	mode := os.Getenv("MSA_MODE")
	if "" == mode {
		mode = "debug"
	}
	conf, err := config.NewConfig()
	if nil != err {
		panic(err)
	}

	if "debug" == mode {
		logger.DefaultLogger = zerolog.NewLogger(
			logger.WithOutput(os.Stdout),
			logger.WithLevel(logger.TraceLevel),
			zerolog.WithDevelopmentMode(),
		)
		logger.Warn("Running in \"debug\" mode. Switch to \"release\" mode in production.")
		logger.Warn("- using env:	export MSA_MODE=release")
	} else {
		logger.DefaultLogger = zerolog.NewLogger(
			logger.WithOutput(os.Stdout),
			logger.WithLevel(logger.TraceLevel),
			zerolog.WithProductionMode(),
		)
	}

	setupEnvironment()

	// load default config
	logger.Infof("default config is: \n\r%v", defaultYAML)
	memorySource := memory.NewSource(
		memory.WithYAML([]byte(defaultYAML)),
	)
	conf.Load(memorySource)
	conf.Scan(&Schema)

	// merge others
	if "file" == configDefine.Source {
		mergeFile(conf)
	} else if "consul" == configDefine.Source {
		mergeConsul(conf)
	} else if "etcd" == configDefine.Source {
		mergeEtcd(conf)
	}

	ycd, err := goYAML.Marshal(&Schema)
	if nil != err {
		logger.Error(err)
	} else {
		logger.Infof("current config is: \n\r%v", string(ycd))
	}

	level, err := logger.GetLevel(Schema.Logger.Level)
	if nil != err {
		logger.Warnf("the level %v is invalid, just use info level", Schema.Logger.Level)
		level = logger.InfoLevel
	}
	logger.Init(
		logger.WithLevel(level),
	)
}
