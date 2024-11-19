package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port   string `yaml:"port"`
	Log    Log    `yaml:"log"`
	Rate   Rate   `yaml:"rate"`
	MySql  MySql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Attach Attach `yaml:"attach"`
}

type Log struct {
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Level      string `yaml:"level"`
}

type Rate struct {
	Limit      int `yaml:"limit"`
	Burst      int `yaml:"burst"`
	ResetTimes int `yaml:"reset_times"`
}

type MySql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Password    string `yaml:"password"`
	Db          int    `yaml:"db"`
	PoolSize    int    `yaml:"pool_size"`
	MinIdleConn int    `yaml:"min_idle_conn"`
}

type Attach struct {
	Blob string `json:"blob"`
}

var config *Config

func GetConfig() *Config {
	return config
}

func InitConfig(rootPath string) error {
	configPath := filepath.Join(rootPath, "/config/config.yaml")
	dataBytes, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	config = &Config{}
	err = yaml.Unmarshal(dataBytes, config)
	if err != nil {
		return err
	}
	return nil
}
