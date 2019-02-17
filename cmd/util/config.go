package util

import (
	"github.com/kelseyhightower/envconfig"
)

type MySQL struct {
	Port     int    `default:"3306"`
	User     string `default:"root"`
	Host     string `default:"localhost"`
	Password string `default:"mysql"`
	Database string `default:"helpnamiki"`
}

type Config struct {
	Host   string `default:"localhost"`
	Port   string `default:"9090"`
	Client string `default:"localhost:8080"`
	MySQL  MySQL  `envconfig:"MYSQL"`
}

var config Config

func (c Config) ServerHost() string {
	return c.Client
}

func init() {
	config, _ = Init()
}

func Init() (Config, error) {
	err := envconfig.Process("NAMIKI", &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func GetConfig() Config {
	return config
}
