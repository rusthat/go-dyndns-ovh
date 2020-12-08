package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

func main() {
	var cfg DynDNSConfig
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Sprintf("Error reading config.yml. Importing environment variables.")
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			fmt.Sprintf("Error reading environment variables. Aborting.")
			processError(err)
		}
	}

	fmt.Sprintf("DNS Zone: %s", cfg.Zone.Name)
	fmt.Sprintf("Name: %s", cfg.Zone.Subdomain)
	fmt.Sprintf("Name: %s", cfg.Zone.USER)
	fmt.Sprintf("Name: %s", cfg.Zone.Password)
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

type DynDNSConfig struct {
	Zone struct {
		Name      string `yaml:"name" env:"OVH_ZONE_NAME" env-description:"dns zone (e.g. domain.com)."`
		Subdomain string `yaml:"subdomain" env:"OVH_ZONE_SUBDOMAIN" env-description:"subdomain whose A record shall be updated."`
		USER      string `yaml:"user" env:"OVH_ZONE_USER" env-description:"user id for the ovh api."`
		Password  string `yaml:"password" env:"OVH_ZONE_PASSWORD" env-description:"The ovh api password"`
	} `yaml:"zone"`
}
