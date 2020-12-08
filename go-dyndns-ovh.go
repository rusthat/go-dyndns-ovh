package main

import (
	"encoding/base64"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func main() {
	var config = getConfig()     // Read config.yml or ENV vars
	checkInternetConn()          // Check if connection to internet is available
	var publicIP = getPublicIP() // Get the current public IP address

	fmt.Println(fmt.Sprintf("OVH Domain DynDNS Update utility\n\nSubdomain: %s\nUser: %s\nPass: %s\nPublic IP: %s\n", config.Zone.Subdomain, config.Zone.USER, config.Zone.PASS, publicIP))
	var res = updateDNS(publicIP, config)
	fmt.Println(res)
	//TODO: Add loop time to config
	//TODO: Check if IP has changed every loop iteration
	//TODO: Update DNS only on change
	//TODO: Check updateDNS response for successfull change
	// response: nochg 255.255.255.255 = adress up to date
}

func updateDNS(pubip string, config DynDNSConfig) string {
	var response string

	req, err := http.NewRequest("GET", fmt.Sprintf("http://www.ovh.com/nic/update?system=dyndns&hostname=%s&myip=%s", config.Zone.Subdomain, pubip), nil)
	if err == nil {
		req.Header.Add("Authorization", "Basic "+basicAuth(config.Zone.USER, config.Zone.PASS))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err == nil {
			// successfull request
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				resp.Body.Close()
				response = string(body)
			} else {
				processError(err)
			}
		}
	} else {
		processError(err)
	}
	return response
}

func getPublicIP() string {
	var ip string
	resp, err := http.Get("http://ipecho.net/plain")
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err == nil {
			ip = string(body)
		} else {
			processError(err)
		}
	} else {
		processError(err)
	}
	return ip
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func getConfig() DynDNSConfig {
	// first try reading configuration from config.yml
	var dnsConfig DynDNSConfig
	var err = cleanenv.ReadConfig("config.yml", &dnsConfig)

	if err != nil {
		// secondly try parsing env vars
		err := cleanenv.ReadEnv(&dnsConfig)
		if err != nil {
			processError(err)
		}
	}

	return dnsConfig
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func checkInternetConn() {
	_, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Connection check failed. Aborting.")
		processError(err)
	}
}

type DynDNSConfig struct {
	Zone struct {
		Subdomain string `yaml:"subdomain" env:"OVH_ZONE_SUBDOMAIN" env-description:"subdomain whose A record shall be updated."`
		USER      string `yaml:"user" env:"OVH_ZONE_USER" env-description:"user id for the ovh api."`
		PASS      string `yaml:"pass" env:"OVH_ZONE_PASS" env-description:"The ovh api password"`
	} `yaml:"zone"`
}
