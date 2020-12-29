package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var configFileCache []byte

type ConfigFile struct {
	Redis  Redis    `yaml:"redis"`
	Config []Config `yaml:"config"`
}
type Config struct {
	Root      string `yaml:"root"`
	Server    string `yaml:"server"`
	Port      int    `yaml:"port"`
	Instances int    `yaml:"instances"`
	Index     string `yaml:"index"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

func GetConfigFile() ([]byte, error) {
	if configFileCache == nil {
		yamlFile, err := ioutil.ReadFile("config.yaml")
		configFileCache = yamlFile
		return yamlFile, err
	} else {
		return configFileCache, nil

	}
}

func (c *ConfigFile) GetConfig() *ConfigFile {

	// yamlFile, err := ioutil.ReadFile("config.yaml")
	yamlFile, err := GetConfigFile()
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
func Parse() *ConfigFile {
	var c ConfigFile
	return c.GetConfig()
}
