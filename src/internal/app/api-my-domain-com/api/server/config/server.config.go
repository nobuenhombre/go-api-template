package config

import (
	"github.com/nobuenhombre/suikat/pkg/fico"
	"github.com/nobuenhombre/suikat/pkg/ge"
	"gopkg.in/yaml.v3"
)

type RedisConfig struct {
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Password string `yaml:"password,omitempty"`
	DB       int    `yaml:"db,omitempty"`
}

type HTTPServerConfig struct {
	Host  string      `yaml:"host,omitempty"`
	Port  string      `yaml:"post,omitempty"`
	Store RedisConfig `yaml:"store,omitempty"`
}

type HostsConfig struct {
	API HTTPServerConfig `yaml:"api,omitempty"`
}

type Config struct {
	Hosts HostsConfig `yaml:"hosts,omitempty"`
}

func (c *Config) Load(fileName string) error {
	txtConfigFile := fico.TxtFile(fileName)

	configData, err := txtConfigFile.Read()
	if err != nil {
		return ge.Pin(err)
	}

	err = yaml.Unmarshal([]byte(configData), c)
	if err != nil {
		return ge.Pin(err)
	}

	return nil
}

func (c *Config) Save(fileName string) error {
	txtConfigFile := fico.TxtFile(fileName)

	configData, err := yaml.Marshal(c)
	if err != nil {
		return ge.Pin(err)
	}

	err = txtConfigFile.Write(string(configData))
	if err != nil {
		return ge.Pin(err)
	}

	return nil
}
