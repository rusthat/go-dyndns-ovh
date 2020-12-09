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

var config Config

func main() {
	checkInternetConn()                                   // Check if connection to internet is available
	var dynDNSConfig = getConfig()                        // Read dynDNSConfig.yml or ENV vars
	var publicIP = httpGetBody("http://ipecho.net/plain") // Get the current public IP address

	fmt.Println(fmt.Sprintf("OVH Domain DynDNS Update utility\n\nDNS Record: %s\nUser: %s\nPass: %s\nPublic IP: %s\n", dynDNSConfig.RECORD, dynDNSConfig.USER, dynDNSConfig.PASS, publicIP))
	var res = updateDNS(publicIP, dynDNSConfig)
	fmt.Println(res)
	//TODO: Add loop time to dynDNSConfig
	//TODO: Check if IP has changed every loop iteration
	//TODO: Update RECORD only on change
	//TODO: Check updateDNS response for sucess
	// response: nochg 255.255.255.255 = address up to date
}

func updateDNS(pubip string, config Config) string {
	var response string

	req, err := http.NewRequest("GET", fmt.Sprintf("http://www.ovh.com/nic/update?system=dyndns&hostname=%s&myip=%s", config.RECORD, pubip), nil)
	if err == nil {
		req.Header.Add("Authorization", "Basic "+basicAuth(config.USER, config.PASS))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err == nil {
			// successfull request
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				err := resp.Body.Close()
				if err == nil {
					response = string(body)
				} else {
					processError(err)
				}
			} else {
				processError(err)
			}
		}
	} else {
		processError(err)
	}
	return response
}

func httpGetBody(url string) string {
	var response string
	resp, err := http.Get(url)
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			// successful request, close connection
			err := resp.Body.Close()
			if err == nil {
				// set ip from body
				response = string(body)
			}
		} else {
			processError(err)
		}
	} else {
		processError(err)
	}
	return response
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func getConfig() Config {
	// First try parsing environment variables
	var dnsConfig Config
	err := cleanenv.ReadEnv(&dnsConfig)
	if err != nil {
		fmt.Print("Error parsing environment variables.")

		// If that fails try reading .env file
		err = cleanenv.ReadConfig(".env", &dnsConfig)
		if err != nil {
			fmt.Print("Error reading .env file.")
			processError(err)
		}
	}
	return dnsConfig
}

func checkConfig(config Config) {
	if (len(config.USER) == 0) ||
		(len(config.RECORD) == 0) ||
		(len(config.PASS) == 0) {
		printHelp(config)
		os.Exit(2)
	}
}

func processError(err error) {
	printHelp(config)
	fmt.Println(err)
	os.Exit(2)
}

func printHelp(config Config) {
	var help, err = cleanenv.GetDescription(config, nil)
	if err == nil {
		fmt.Print(help)
	} else {
		processError(err)
	}
}

func checkInternetConn() {
	_, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Connection check failed. Aborting.")
		processError(err)
	}
}

type Config struct {
	RECORD string `env:"OVH_DNS_RECORD" env-description:"subdomain whose A record shall be updated."`
	USER   string `env:"OVH_DNS_USER" env-description:"user id for the ovh api."`
	PASS   string `env:"OVH_DNS_PASS" env-description:"The ovh api password"`
	LOOP   int    `env:"OVH_DNS_LOOP" env-default:"0" env-description:"The ovh api password"`
}
