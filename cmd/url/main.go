package main

import (
	"fmt"
	app "github.com/NautiloosGo/url_shortener/internal/app"
	st "github.com/NautiloosGo/url_shortener/internal/storage"
	"strconv"
	"time"
)

func main() {

	Conf := st.LoadConfiguration("./config.json")
	Catalog := st.LoadDB(Conf.FileCatalog)
	fmt.Println(Conf)
	fmt.Println(Catalog)
	fmt.Println(fmt.Sprintf(GetRandomString(setup.Settings.Qty, setup.Settings.Letters)))

}
