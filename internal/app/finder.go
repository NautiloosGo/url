package app

import (
	st "github.com/NautiloosGo/url/internal/storage"
)

// find url by short url in catalog
func FindSurl(cat st.Catalog, url string) (string, bool) {
	for _, c := range cat.List {
		if c.Surl == url {
			return c.Url, true
		}
	}
	return "", false
}

// find short url by url in catalog
func FindUrl(cat st.Catalog, url string) (string, bool) {
	for _, c := range cat.List {
		if c.Url == url {
			return c.Surl, true
		}
	}
	return "", false
}
