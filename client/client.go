package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"time"
)

const Domain = "http://127.0.0.1:8080/"

type urlList struct {
	List      []string
	Listshort []string
}

const filedb = "./client/urls.json"

func LoadURLFromDB(file string) urlList {
	var db = urlList{}
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(readFile)
	jsonParser.Decode(&db)
	return db
}

// simple GET
func Get(surl string) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", Domain, nil)
	if err != nil {
		log.Println(err)
	}
	params := request.URL.Query()
	params.Add("url", "")
	params.Add("short_url", surl)
	request.URL.RawQuery = params.Encode()
	// send
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	// response
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(result))
}

// simple POST
func Post(url string) {
	client := &http.Client{}
	request, err := http.NewRequest("POST", Domain, nil)
	if err != nil {
		log.Println(err)
	}
	params := request.URL.Query()
	params.Add("url", url)
	params.Add("short_url", "")
	request.URL.RawQuery = params.Encode()
	// send
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	// response
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(result))
}

// POST as JSON
func PostJson(url string) {
	var std map[string]string = map[string]string{"url": url, "short_url": ""}
	data, err := json.Marshal(std)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", Domain, body)
	if err != nil {
		log.Println(err)
	}
	// header
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(result))
}

func main() {
	// get urls from .json
	db := LoadURLFromDB(filedb)
	for i, url := range db.List {
		switch i % 2 {
		case 0:
			Post(url)
			PostJson(url)
		case 1:
			PostJson(url)
			Post(url)
		}
	}
	for _, surl := range db.Listshort {
		Get(surl)
	}
}
