package main

import (
	"encoding/json"
	"fmt"
	app "github.com/NautiloosGo/url/internal/app"
	st "github.com/NautiloosGo/url/internal/storage"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Conf st.Config
var Catalog st.Catalog

type responseToClient struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	var ok = "success"
	// client
	fmt.Println("client connect success ", r.RemoteAddr)
	// request
	fmt.Println(r.Method, r.RequestURI)
	data := make(map[string]string)
	// read
	buf := make([]byte, 2048)
	n, _ := r.Body.Read(buf)
	// operating
	if r.Method == "GET" {
		r.ParseForm()
		for k, v := range r.Form {
			if k == "short_url" {
				url, found := FindSurl(v[0])
				if found == true {
					data["url"] = url
					break
				} else {
					data["url"] = ""
					data["short_url"] = v[0]
					ok = "not found"
				}

			}

		}
	}

	if r.Method == "POST" {
		ct, k := r.Header["Content-Type"]
		if k {
			// check for json
			if ct[0] == "application/json" {
				json.Unmarshal(buf[:n], &data)
			}
		} else {
			r.ParseForm()
			for w, v := range r.Form {
				data[w] = v[0]
			}
		}
		url, k := data["url"]
		if k {
			surl, found := FindUrl(url)
			if found == true {
				data["short_url"] = surl
				ok = "short URL already exists"
			} else {
				surl = app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)
				_, found := FindUrl(surl)
				for found == true {
					surl = app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)
					_, found = FindUrl(surl)
				}
				AddLink(url, surl)
				data["short_url"] = surl
				ok = "done"
			}
		}

	}
	// sending
	m := responseToClient{200, ok, data}
	mjson, e := json.Marshal(m)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Fprintf(w, "%v\n", string(mjson))
}

func UploadCatalog() {
	Catalog = st.LoadDB(Conf.FileCatalog)
	fmt.Println("Success: Catalog from ", Conf.FileCatalog)
}

func FindSurl(url string) (string, bool) {
	for _, c := range Catalog.List {
		if c.Surl == url {
			return c.Url, true
		}
	}
	return "", false
}

func FindUrl(url string) (string, bool) {
	for _, c := range Catalog.List {
		if c.Url == url {
			return c.Surl, true
		}
	}
	return "", false
}

func AddLink(url, surl string) {
	req := st.Request{
		Id:   "",
		Url:  url,
		Surl: surl,
	}
	Catalog.List = append(Catalog.List, req)
}

func AutosaverDB(filedb string, n time.Duration) {
	for {
		<-time.After(n)
		//back in .json
		rawDataOut, err := json.MarshalIndent(&Catalog, "", "  ")
		if err != nil {
			fmt.Println("JSON marshaling failed:", err)
		}

		err = ioutil.WriteFile(filedb, rawDataOut, 0)
		if err != nil {
			fmt.Println("Cannot write updated catalog file:", err)
		}
	}
}

func main() {
	//get config
	Conf = st.LoadConfiguration("./config.json")
	fmt.Println("Configs uploaded")
	//get local db (json)
	UploadCatalog()
	//autosave every n seconds
	go AutosaverDB(Conf.FileCatalog, time.Second*5)
	//start server
	http.HandleFunc("/", defaultFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
