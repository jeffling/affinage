package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var AppConfig *config

type config struct {
	StripeKey string `json "stripeKey"`
	Port      string `json "port"`
}

func init() {
	// get config
	configPath := "./conf.json"
	configString, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Could not find config file", err)
	}
	AppConfig = &config{}
	if err := json.Unmarshal(configString, &AppConfig); err != nil {
		log.Fatal("Error parsing config: ", err)
	}
}
