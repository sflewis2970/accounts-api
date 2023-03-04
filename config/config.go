package config

import (
	"log"
	"os"
)

// Config variable keys
const (
	// ENV System ENV setting
	ENV string = "ENV"

	// HOST system info
	HOST string = "HOST"
	PORT string = "PORT"

	POSTGRES_HOST string = "postgres_host"
	POSTGRES_PORT string = "postgres_port"
	POSTGRES_USER string = "postgres_user"
)

// PRODUCTION Config variable values
const (
	PRODUCTION string = "PROD"
)

type PostGreSQL struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
}

type CfgData struct {
	Env        string `json:"env"`
	Host       string `json:"hostname"`
	Port       string `json:"hostport"`
	PostGreSQL PostGreSQL
}

type Config struct {
	cfgData *CfgData
}

var config *Config

// Unexported type functions
func (c *Config) loadConfigEnv() {
	// Loading config environment variables
	log.Print("loading config environment variables...")

	// Load host config data
	c.cfgData.Env = os.Getenv(ENV)
	c.cfgData.Host = os.Getenv(HOST)
	c.cfgData.Port = os.Getenv(PORT)

	// Load redis config data
	c.cfgData.PostGreSQL.Host = os.Getenv(POSTGRES_HOST)
	c.cfgData.PostGreSQL.Port = os.Getenv(POSTGRES_PORT)
	c.cfgData.PostGreSQL.User = os.Getenv(POSTGRES_USER)
}

func (c *Config) LoadCfgData() *CfgData {
	log.Print("Using config environment to load config")

	c.loadConfigEnv()

	return c.cfgData
}

func NewConfig() *Config {
	if config == nil {
		log.Print("creating config object")

		// Initialize config
		config = new(Config)

		// Initialize config data
		config.cfgData = new(CfgData)
	} else {
		log.Print("returning config object")
	}

	return config
}
