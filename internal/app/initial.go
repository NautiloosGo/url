package app

import (
	"fmt"
	st "github.com/NautiloosGo/url/internal/storage"
)

var Conf st.Config
var Catalog st.Catalog

// upload configurations and catalog
func Initial() error {
	//get config
	Conf = st.LoadConfiguration(st.GetConfAdr())
	fmt.Println("Configs uploaded from: ", st.GetConfAdr())
	if Conf.DBtype == "local" {
		//get local db (json)
		Catalog = st.UploadCatalog(Conf.FileCatalog)
	} else {
		fmt.Println("DB not found")
	}
	return nil
}

func GetCatalog() *st.Catalog {
	return &Catalog
}

func GetSaveTimer() int {
	return Conf.AutosaveTimer
}

func GetDBtype() string {
	return Conf.DBtype
}
