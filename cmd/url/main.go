package main

import (
	"encoding/json"
	"fmt"
	app "github.com/NautiloosGo/url/internal/app"
	st "github.com/NautiloosGo/url/internal/storage"
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
	//check header
	ct, k := r.Header["Content-Type"]
	if k {
		// check for json
		if ct[0] == "application/json" {
			buf := make([]byte, 2048)
			n, _ := r.Body.Read(buf)
			json.Unmarshal(buf[:n], &data)
		}
		// todo other types
	} else {
		r.ParseForm()
		for w, v := range r.Form {
			data[w] = v[0]
		}
	}
	var NewLink = st.Request{}
	NewLink.Url = data["url"]
	NewLink.Surl = data["short_url"]

	NewLink, ok = MetodSwitcher(NewLink, r)

	// // operating
	// if r.Method == "GET" {
	// 	r.ParseForm()
	// 	for k, v := range r.Form {
	// 		if k == "short_url" {
	// 			url, found := app.FindSurl(Catalog, v[0])
	// 			if found == true {
	// 				data["url"] = url
	// 				break
	// 			} else {
	// 				data["url"] = ""
	// 				data["short_url"] = v[0]
	// 				ok = "not found"
	// 			}

	// 		}

	// 	}
	// }

	// if r.Method == "POST" {
	// 	ct, k := r.Header["Content-Type"]
	// 	if k {
	// 		// check for json
	// 		if ct[0] == "application/json" {
	// 			json.Unmarshal(buf[:n], &data)
	// 		}
	// 	} else {
	// 		r.ParseForm()
	// 		for w, v := range r.Form {
	// 			data[w] = v[0]
	// 		}
	// 	}
	// 	url, k := data["url"]
	// 	if k {
	// 		surl, found := app.FindUrl(Catalog, url)
	// 		if found == true {
	// 			data["short_url"] = surl
	// 			ok = "short URL already exists"
	// 		} else {
	// 			surl = app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)
	// 			_, found := app.FindUrl(Catalog, surl)
	// 			for found == true {
	// 				surl = app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)
	// 				_, found = app.FindUrl(Catalog, surl)
	// 			}
	// 			AddLink(url, surl)
	// 			data["short_url"] = surl
	// 			ok = "done"
	// 		}
	// 	}

	// }

	// sending
	data["url"] = NewLink.Url
	data["short_url"] = NewLink.Surl
	m := responseToClient{200, ok, data}
	mjson, e := json.Marshal(m)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Fprintf(w, "%v\n", string(mjson))
}

func MetodSwitcher(data st.Request, r *http.Request) (st.Request, string) {
	if r.Method == "GET" {
		return Get(data)
	}
	if r.Method == "POST" {
		return Post(data)
	}
	return data, "unknown Method"
}

func Get(data st.Request) (st.Request, string) {
	if data.Surl != "" {
		url, found := app.FindSurl(Catalog, data.Surl)
		if found {
			data.Url = url
			return data, "done"
		} else {
			data.Url = ""
			return data, "url not found"
		}
	} else {
		return data, "requested short_url is empty"
	}
}

func Post(data st.Request) (st.Request, string) {
	fmt.Println("data", data)
	if data.Url != "" {
		fmt.Println("befor search", data)
		surl, found := app.FindUrl(Catalog, data.Url)
		if found {
			fmt.Println("if found", data)
			data.Surl = surl
			return data, "done. already exists"
		} else {
			fmt.Println("not foound", data)
			return PostUniq(data)

			// surl = app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)
			// if _, found := app.FindUrl(Catalog, surl); found {

			// }
			// for found == true {
			// 	surl = app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)
			// 	_, found = app.FindUrl(Catalog, surl)
			// }
			// AddLink(url, surl)
			// data["short_url"] = surl
			// ok = "done"
		}
	} else {
		return data, "requested url is empty"
	}
}

func PostUniq(data st.Request) (st.Request, string) {
	fmt.Println("uniq", data)
	surl := app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)
	if _, found := app.FindSurl(Catalog, surl); found {
		fmt.Println("uniq WOW", data)
		return PostUniq(data)
	} else {
		data.Surl = surl
		fmt.Println("uniq surl", data)
		AddLink(data)
		return data, "done. new short_url"
	}
}

func AddLink(data st.Request) {
	Catalog.List = append(Catalog.List, data)
}

func main() {
	//get config
	Conf = st.LoadConfiguration("./config.json")
	//get local db (json)
	Catalog = st.UploadCatalog(Conf.FileCatalog)
	//autosave every n seconds
	go st.AutosaverDB(Catalog, Conf.FileCatalog, time.Second*5)
	//start server
	http.HandleFunc("/", defaultFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
