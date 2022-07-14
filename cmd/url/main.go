package main

import (
	"fmt"
	app "github.com/NautiloosGo/url/internal/app"
	st "github.com/NautiloosGo/url/internal/storage"
)

func main() {

	Conf := st.LoadConfiguration("./config.json")
	Catalog := st.LoadDB(Conf.FileCatalog)
	fmt.Println(Conf)
	fmt.Println(Catalog)
	fmt.Println(fmt.Sprintf(app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)))

}
