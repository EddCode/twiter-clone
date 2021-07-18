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

type config struct {
	Database database `yaml:"database"`
	Server   server   `yaml:"server"`
}

func GetConfig(configFile string) (*config, error) {
    file, err := os.Open(configFile)

    if err != nil {
        return nil, err
    }

    defer file.Close()

    setting := &config{}
    yd := yaml.NewDecoder(file)
    err = yd.Decode(setting)

    if err != nil {
        return nil, err
    }

    return setting, nil
}
