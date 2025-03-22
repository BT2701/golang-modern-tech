package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
}

// DatabaseConfig struct
type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

// ServerConfig struct
type ServerConfig struct {
	Port string
}

// JWTConfig struct
type JWTConfig struct {
	Key string
}
