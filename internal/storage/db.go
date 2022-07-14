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

// func AutosaverDB(filedb string, n int) {
// 	for {
// 		<-time.After(time.Second * n)
// 		//back in .json
// 		rawDataOut, err := json.MarshalIndent(&Catalog, "", "  ")
// 		if err != nil {
// 			fmt.Println("JSON marshaling failed:", err)
// 		}

// 		err = ioutil.WriteFile(filedb, rawDataOut, 0)
// 		if err != nil {
// 			fmt.Println("Cannot write updated catalog file:", err)
// 		}
// 	}
// }
