package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const configYML = "config.yml"

var cfg *Config

func ConfigInstance() *Config {
	return cfg
}

// Database - сontains all parameters database connection
type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"database"`
	SslMode  string `yaml:"sslmode"`
	Driver   string `yaml:"driver"`
}

// Grpc - сontains parameter address grpc
type Grpc struct {
	Address string `yaml:"address"`
}

// Json - сontains parameter rest json connection
type Json struct {
	Address string `yaml:"address"`
}

// Project - сontains all parameters project information
type Project struct {
	Name    string `yaml:"name"`
	Author  string `yaml:"author"`
	Version string `yaml:"version"`
}

// Config - сontains all configuration parameters
type Config struct {
	Project  Project  `yaml:"project"`
	Grpc     Grpc     `yaml:"grpc"`
	Json     Json     `yaml:"json"`
	Database Database `yaml:"database"`
}

// ReadConfigYML - read configurations from file and init instance Config
func ReadConfigYML() error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}
