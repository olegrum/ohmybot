package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config struct to save service configuration
type Config struct {
	Main struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
	} `yaml:"main"`
	AdminPanel struct {
		Endpoint string `yaml:"endpoint"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
	} `yaml:"admin_panel"`
}

//ReadConfig function to read config file
func ReadConfig(file string) (Config, error) {
	yfile, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err2 := yaml.Unmarshal(yfile, &config)
	if err2 != nil {
		return Config{}, err2
	}
	return config, nil

}
