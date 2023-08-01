package config

import (
	"github.com/caarlos0/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServerConfig struct {
	Environment string `env:"ENV,required"`
	Port        string `env:"PORT,required"`
}

func LoadEnvConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func NewDB() *gorm.DB {
	dsn := "user:password@tcp(mysql)/db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
