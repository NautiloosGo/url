package app

import (
	"fmt"
	st "github.com/NautiloosGo/url/internal/storage"
	"time"
)

var Conf st.Config
var Catalog st.Catalog

func Initial() error {
	//get config
	Conf = st.LoadConfiguration("./config.json")
	fmt.Println("Configs uploaded")
	//get local db (json)
	Catalog = st.UploadCatalog(Conf.FileCatalog)
	//autosave every n seconds
	go st.AutosaverDB(Catalog, Conf.FileCatalog, time.Second*5)
	//start server
	return nil
}
