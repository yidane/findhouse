package config

import (
	"io/ioutil"
)

type Config struct {
	globalContent     map[string]string
	selectionContents map[string]map[string]string
	sections          []string
}

func NewConfig() *Config {
	return &Config{
		globalContent:     make(map[string]string),
		selectionContents: make(map[string]map[string]string),
	}
}

func (p *Config) Load(configFilePath string) *Config {
	stream, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic("read config file error :" + configFilePath + "\n")
	}
	p.LoadString(string(stream))
	return p
}

func (p *Config) LoadString(configString string) error {
	return nil
}
