package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const BC = "http://127.0.0.1:8080/"

// Отправить простой запрос HTTP GET
func httpSimpleGet() {
	resp, err := http.Get(BC + "index?aa=AA&bb=BB")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	// Получить ответный контент
	resultByte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(resultByte))
}

// Получить запрос, который устанавливает заголовок запроса и параметры запроса
func httpGet() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", BC, nil)
	if err != nil {
		log.Println(err)
	}
	// Добавить пользовательские данные в заголовок запроса
	request.Header.Add("client", "T")
	// Добавить параметры запроса
	params := request.URL.Query()
	params.Add("short_url", "abc123123d")
	params.Add("url", "ahggg")
	request.URL.RawQuery = params.Encode()
	// Отправляем http запрос, запрос успешен, получаем ответ
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	// Получить весь контент ответа
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// выводим содержимое ответа
	fmt.Println(string(result))
}

// HTTP POST запрос и отправка данных JSON
func httpPostJson() {
	// Опубликовать данные JSON широко используются
	// Отправляем данные json, мы обычно используем карту или структуру для хранения данных
	// Затем преобразовать в данные JSON
	// Затем преобразуем его в байтовые данные и отправляем в теле отправки
	// Давайте смоделируем этот процесс
	var std map[string]string = map[string]string{"work": "programmer", "skills": "golang", "addr": "Пекин"}
	data, err := json.Marshal(std)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", BC, body)
	if err != nil {
		log.Println(err)
	}
	// Установить заголовок запроса
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(result))
}

// Имитация публикации для отправки данных формы
func HttpPosForm() {
	formData := url.Values{}
	formData.Set("userName", "admin")
	formData.Set("userPwd", "admin123456")
	req, err := http.NewRequest("POST", BC, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(formData.Encode())))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))

}

// Имитация клиента для отправки запроса PUT
func HttpPut() {
	// Отправить запрос PUT и POST запрос может отправить данные JSON и формы
	std := map[string]string{"method": "PUT"}
	data, err := json.Marshal(std)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("PUT", BC, body)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(result))
}

// смоделированный клиент отправляет запрос DELETE
func HttpDelete() {
	req, err := http.NewRequest("DELETE", BC, nil)
	if err != nil {
		log.Println(err)
	}
	// Добавить параметр запроса аналогично отправке запроса получения
	params := req.URL.Query()
	params.Add("user", "pahnaskdjalsdklasd")
	req.URL.RawQuery = params.Encode()
	// Отправляем http запрос, запрос успешен, получаем ответ
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	// Получить весь контент ответа
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// выводим содержимое ответа
	fmt.Println(string(result))
}
func main() {

	//httpSimpleGet()
	httpGet()
	//httpPostJson()
	//HttpPosForm()
	//HttpPut()
	//HttpDelete()
}
