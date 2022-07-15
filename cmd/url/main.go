package main

import (
	app "github.com/NautiloosGo/url/internal/app"
	serv "github.com/NautiloosGo/url/internal/server"
)

func main() {
	// upload cofig and db
	app.Initial()
	// start server
	serv.StartServe()
}
