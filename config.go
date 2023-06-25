package main

import (
	"io"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server  string    `toml:"server"`
	Token   string    `toml:"token"`
	Webhook []Webhook `toml:"webhook"`
}

func LoadConfig(r io.Reader) (*Config, error) {
	var config Config

	d := toml.NewDecoder(r)
	err := d.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
