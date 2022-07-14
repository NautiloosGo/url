package main

import (
	"encoding/json"
	"fmt"
	//app "github.com/NautiloosGo/url/internal/app"
	st "github.com/NautiloosGo/url/internal/storage"
	"log"
	"net/http"
	"time"
)

var Conf st.Config
var Catalog st.Catalog

type responseToClient struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	var dataForm map[string][]string
	var ok = "success"
	// Идентифицируем клиентское соединение
	fmt.Println("client connect success ", r.RemoteAddr)
	// Получить содержимое адресной строки
	fmt.Println(r.Method, r.RequestURI)
	// Получить содержимое заголовка запроса
	// for k, v := range r.Header {
	// 	fmt.Println(k, v[0])
	// }
	data := make(map[string]string)
	if err := r.ParseForm(); err == nil {
		if r.Form != nil {
			dataForm = r.Form
		}
	}
	// Читаем содержимое клиента
	buf := make([]byte, 2048)
	n, _ := r.Body.Read(buf)
	// Получить содержимое в теле запроса
	fmt.Println("receive data from body", string(buf[:n]))

	if r.Method == "GET" {
		r.ParseForm()
		for k, v := range r.Form {
			if k == "short_url" {
				url, found := FindSurl(v[0])
				if found == true {
					data["url"] = url
					break
				} else {
					data["url"] = ""
					ok = "unsuccess"
				}

			}

		}
	}
	// Обработка запросов POST и PUT, отправленных клиентом
	if r.Method == "POST" || r.Method == "PUT" {
		ct, ok := r.Header["Content-Type"]
		if ok {
			// Если это данные json, судите по заголовку запроса
			if ct[0] == "application/json" {
				json.Unmarshal(buf[:n], &data)
			}
			// Если это данные формы POST
			if ct[0] == "application/x-www-form-urlencoded" {
				if dataForm != nil {
					for k, v := range dataForm {
						data[k] = v[0]
					}
				}
			}
		}
	}
	// Обработка запроса DELETE клиента
	// Запись текущего времени `2006-01-02 15: 04: 05` относится к формату формата
	data["time"] = time.Now().Format("2006-01-02 15:04:05")
	m := responseToClient{200, ok, data}
	mjson, e := json.Marshal(m)
	if e != nil {
		fmt.Println(e)
	}
	// ответить клиенту в формате json
	fmt.Fprintf(w, "%v\n", string(mjson))
}
func main() {
	Start()

	http.HandleFunc("/", defaultFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func Start() {
	Conf = st.LoadConfiguration("./config.json")
	Catalog = st.LoadDB(Conf.FileCatalog)
	fmt.Println("Configuration: ", Conf)
	fmt.Println("Catalog: ", Catalog)
	//fmt.Println(fmt.Sprintf(app.GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)))

}

func FindSurl(url string) (string, bool) {
	fmt.Println("get", url)
	for _, c := range Catalog.List {
		if c.Surl == url {
			return c.Url, true
		}
	}
	return "", false
}
