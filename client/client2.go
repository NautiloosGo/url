package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const Domain = "http://127.0.0.1:8080/"

type urlList struct {
	List      []string
	Listshort []string
}

const surl = "secondclientGet"

// simple GET
func SecondGet(surl string) {
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
	fmt.Println("SECOND CL REQ: ", request)
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

func main() {
	surl := "secondclientGet"
	for i := 0; i <= 20; i++ {
		SecondGet(surl)
	}
}
