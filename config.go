package main

import (
	"fmt"
	toml "github.com/pelletier/go-toml"
	"io/ioutil"
)

/* Global config file */
var ConfigFile Config

type Config struct {
	Instance     map[string]Instance
	AwsAccounts	 map[string]AwsAccounts
}

type Instance struct {
	InstanceId string
	Region     string
	AwsAccount string
}

type AwsAccounts struct {
	AwsKeyId     string
	AwsSecretKey string
	AwsToken     string
}

func LoadConfig() {

	/* Read config file */
	buffer, err := ioutil.ReadFile("config.toml")
	if err != nil {
		fmt.Println("Could not open config file")
		panic(err)
	}

	/* Map config file */
	if err := toml.Unmarshal(buffer, &ConfigFile); err != nil {
		fmt.Println("Error Unmarshal file")
		panic(err)
	}

}
