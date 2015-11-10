package configLoader

import (
	"os"
	"io/ioutil"
)


type ConfigLoader struct {
	configs map[string]*Config
}

func NewConfigLoader() *ConfigLoader{
	configLoader := ConfigLoader{configs: make(map[string]*Config)}
	return &configLoader
}

func (this *ConfigLoader) GetConfigFor(file string) (*Config, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return NewConfig(), err
	}
	if cachedConfig, ok := this.configs[file]; ok {
		return cachedConfig, nil
	}

	this.configs[file] = NewConfig()

	fileContents, err := ioutil.ReadFile(file)

	if err != nil {
		return NewConfig(), err
	}

	this.configs[file].Parse(fileContents)

	return this.configs[file], nil
}
