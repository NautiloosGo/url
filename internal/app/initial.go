package app

import (
	"fmt"
	st "github.com/NautiloosGo/url/internal/storage"
)

var Conf st.Config
var Catalog st.Catalog

func Initial() error {
	//get config
	Conf = st.LoadConfiguration(st.GetConfAdr())
	fmt.Println("Configs uploaded from: ", st.GetConfAdr())
	//get local db (json)
	Catalog = st.UploadCatalog(Conf.FileCatalog)
	return nil
}
func GetCatalog() *st.Catalog {
	return &Catalog
}
