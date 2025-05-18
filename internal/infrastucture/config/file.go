package config

import "github.com/caarlos0/env/v6"

type Config struct {
	WebServer struct {
		RESTAPI struct {
			Address string `env:"WEBSERVER_RESTAPI_ADDRESS" envDefault:":8080"`
		}
	}
	Logger struct {
	}
}

func NewConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) GetServerHTTPAddress() string {
	return c.WebServer.RESTAPI.Address
}
