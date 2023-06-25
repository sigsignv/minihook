package main

import (
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server  string    `toml:"server"`
	Token   string    `toml:"token"`
	Webhook []Webhook `toml:"webhook"`
}

func NewConfig(r io.Reader) (*Config, error) {
	var config Config

	d := toml.NewDecoder(r)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return NewConfig(f)
}
