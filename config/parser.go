package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ConfigFile struct {
	Redis  Redis    `yaml:"redis"`
	Config []Config `yaml:"config"`
}
type Config struct {
	Root   string `yaml:"root"`
	Server string `yaml:"server"`
	Port   int    `yaml:"port"`
	Index  string `yaml:"index"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

func (c *ConfigFile) GetConfig() *ConfigFile {

	yamlFile, err := ioutil.ReadFile("config.yaml")
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
