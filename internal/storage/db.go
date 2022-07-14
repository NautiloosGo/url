package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func LoadDB(file string) Catalog {
	var catalog Catalog
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(readFile)
	jsonParser.Decode(&catalog)
	return catalog
}
