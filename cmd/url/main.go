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

	if app.GetDBtype() == "local" {
		//autosave every n seconds
		go st.AutosaverDB(app.GetCatalog(), time.Millisecond*time.Duration(app.GetSaveTimer()))
	} else {
		fmt.Println("Postgres status: ", st.Pingdb())
	}
	// start server
	serv.StartServe()
}
