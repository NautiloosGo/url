package main

import (
	"fmt"
	app "github.com/NautiloosGo/url/internal/app"
	serv "github.com/NautiloosGo/url/internal/server"
	st "github.com/NautiloosGo/url/internal/storage"
	"time"
)

func main() {
	// upload cofig and db
	app.Initial()
	for i := 0; i <= 10000; i++ {
		surl := app.GetRandomStringFaster(20, `1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&'()*+,-.:;<=>?@"[]^_{|}~`)
		fmt.Println(`"`, surl, `",`)
	}

	e := false
	if e {
		//autosave every n seconds
		go st.AutosaverDB(app.GetCatalog(), time.Millisecond*time.Duration(app.GetSaveTimer()))
		// start server

		serv.StartServe()
	}
}
