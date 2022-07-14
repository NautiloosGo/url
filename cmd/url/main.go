package main

import (
	"fmt"
	app "github.com/NautiloosGo/url/internal/app"
	st "github.com/NautiloosGo/url/internal/storage"
	"log"
	"net/http"
)

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping", r.RemoteAddr)
	// ответ в формате json
	fmt.Fprintf(w, "%v\n", "answer")

}
func main() {
	http.HandleFunc("/", defaultFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func Start() {
	Conf := st.LoadConfiguration("./config.json")
	Catalog := st.LoadDB(Conf.FileCatalog)
	fmt.Println(Conf)
	fmt.Println(Catalog)
	fmt.Println(fmt.Sprintf(app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)))

}
