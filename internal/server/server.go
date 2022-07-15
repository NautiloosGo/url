package server

import (
	"encoding/json"
	"fmt"
	app "github.com/NautiloosGo/url/internal/app"
	st "github.com/NautiloosGo/url/internal/storage"
	"log"
	"net/http"
)

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
		return app.Get(data)
	}
	if r.Method == "POST" {
		return app.Post(data)
	}
	return data, "unknown Method"
}

func StartServe() {
	//start server
	http.HandleFunc("/", defaultFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}