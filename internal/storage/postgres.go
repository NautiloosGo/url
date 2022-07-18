package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

func GetPostgres(surl string) (Request, string) {
	//surl := "q2ses93yun"
	var req = Request{
		Id:   "",
		Url:  "",
		Surl: "",
	}
	q := `
	select url, surl from UrlShortener where surl = ?; 
		`
	rows, err := db.Query(q, surl)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&req.Url, &req.Surl); err != nil {
			return req, "not found"
		}
		if req.Surl != surl {
			return req, "error postgres request"
		}
		return req, "done. Url from postgres"
	}
	return req, "not found"
}
func PostPostgres(url string) (Request, string) {
	//url := "wqevwr6lkmyunb64et67q2ses93yun"
	var req = Request{
		Id:   "",
		Url:  "",
		Surl: "",
	}
	q := `
	select url, surl from UrlShortener where url = ?; 
		`
	rows, err := db.Query(q, url)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&req.Url, &req.Surl); err != nil {
			return req, "not found"
		}
		if req.Surl != url {
			return req, "error postgres request"
		}
		return req, "done. Url from postgres"
	}
	return req, "not found"
}

func Pingdb() string {
	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	status := "up"
	if err := db.PingContext(ctx); err != nil {
		status = "down"
	}
	return status
}
