package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

// EnvConfig struct
type EnvConfig struct {
	DB struct {
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	}
}

// Vars configuration
func (env EnvConfig) Vars() (EnvConfig, error) {
	f, err := os.Open("env.yaml")
	if err != nil {
		return env, err
	}

	var e EnvConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&e)
	if err != nil {
		return e, err
	}

	return e, nil
}
