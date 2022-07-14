package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
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

func UploadCatalog(file string) Catalog {
	var catalog Catalog
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(readFile)
	jsonParser.Decode(&catalog)
	fmt.Println("Success: Catalog from ", file)
	return catalog
}

func AutosaverDB(c Catalog, filedb string, n time.Duration) {
	for {
		<-time.After(n)
		//back in .json
		rawDataOut, err := json.MarshalIndent(&c, "", "  ")
		if err != nil {
			fmt.Println("JSON marshaling failed:", err)
		}

		err = ioutil.WriteFile(filedb, rawDataOut, 0)
		if err != nil {
			fmt.Println("Cannot write updated catalog file:", err)
		}
	}
}
