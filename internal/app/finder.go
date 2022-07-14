package app

import (
	st "github.com/NautiloosGo/url/internal/storage"
)

func FindSurl(cat st.Catalog, url string) (string, bool) {
	for _, c := range cat.List {
		if c.Surl == url {
			return c.Url, true
		}
	}
	return "", false
}

func FindUrl(cat st.Catalog, url string) (string, bool) {
	for _, c := range cat.List {
		if c.Url == url {
			return c.Surl, true
		}
	}
	return "", false
}

func AddLink(cat st.Catalog, url, surl string) {
	req := st.Request{
		Id:   "",
		Url:  url,
		Surl: surl,
	}
	cat.List = append(cat.List, req)
}
