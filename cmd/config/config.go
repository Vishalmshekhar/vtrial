package config

import (
	"log"
	"os"
    "gopkg.in/yaml.v3"
)
type Configure struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		Version string `yaml:"version"`
	} `yaml:"server"`
	Database struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		Protocol   string `yaml:"protocol"`
		DBName     string `yaml:"dbname"`
		Collection string `yaml:"collection"`
	} `yaml:"database"`
}
var Config Configure
func LoadConfig() Configure {

	data, err := os.ReadFile("./cmd/config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)		
	}
	
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %s", err)		
	}
	return Config;
}