package config

import (
	"github.com/kataras/iris/v12"
	"sync"
)

const (
	Config = "config/config.yaml"
)

var (
	ConfigHandle = &iris.Configuration{}
	Lock = sync.RWMutex{}
)

func InitIrisConfig() {
	//ConfigHandle = iris.WithConfiguration(iris.YAML(Config))
	c := iris.YAML(Config)
	ConfigHandle = &c
}

func GetConfig() *iris.Configuration {
	Lock.RLock()
	defer Lock.RUnlock()
	return ConfigHandle
}