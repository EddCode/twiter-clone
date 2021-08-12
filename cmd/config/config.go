package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type database struct {
	URL     string `yaml:"url"`
	DB      string `yaml:"db"`
	Timeout int    `yaml:"timeout"`
}

type server struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type token struct {
	Secret string `yaml:"secret"`
}

type config struct {
	Database database `yaml:"database"`
	Server   server   `yaml:"server"`
	Token    token    `yaml:"token"`
}

var setting *config

func GetConfig() (*config, error) {
	if setting == nil {
		file, err := os.Open("cmd/config/config.yml")

		if err != nil {
			return nil, err
		}

		defer file.Close()

		setting = &config{}
		yd := yaml.NewDecoder(file)
		err = yd.Decode(setting)

		if err != nil {
			return nil, err
		}

		return setting, nil
	}

	return setting, nil
}
