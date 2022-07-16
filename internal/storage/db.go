package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var catalog Catalog
var config Config

func GetConfAdr() string {
	return confFileAdr
}

func LoadConfiguration(file string) Config {
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
		config = defaultConfig
		fmt.Println(fmt.Sprintf("Invalid catalogs format %s. Using default:\n", file, config, err))
		restoreConfig(defaultConfig, file)
		return config
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		config = defaultConfig
		fmt.Println(fmt.Sprintf("Invalid catalogs format %s. Using default:\n", file, config, err))
		restoreConfig(defaultConfig, file)
		return config
	}
	return config
}

func UploadCatalog(file string) Catalog {
	catalogFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}

		catalog = emptyCatalog
		return catalog
	}
	defer catalogFile.Close()
	jsonParser := json.NewDecoder(catalogFile)
	if err := jsonParser.Decode(&catalog); err != nil {

	}
	fmt.Println("Catalog uploaded from: ", file)
	return catalog
}

func AutosaverDB(c *Catalog, n time.Duration) {
	for {
		<-time.After(n)
		//back in .json

		rawDataOut, err := json.MarshalIndent(&c, "", "  ")
		if err != nil {
			fmt.Println("JSON marshaling failed:", err)
		}

		err = ioutil.WriteFile(config.FileCatalog, rawDataOut, 0)
		if err != nil {
			fmt.Println("Cannot write updated catalog file:", err)
		}
		fmt.Println("Autosaved in ", config.FileCatalog)
	}
}

func restoreConfig(c Config, file string) {
	//back in config.json
	rawDataOut, err := json.MarshalIndent(&c, "", "  ")
	if err != nil {
		fmt.Println("JSON config marshaling failed:", err)
	}

	err = ioutil.WriteFile(file, rawDataOut, 0)
	if err != nil {
		fmt.Println("Cannot write default config file:", err)
	}
}
