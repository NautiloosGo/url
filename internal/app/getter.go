package app

import (
	st "github.com/NautiloosGo/url/internal/storage"
)

func Get(data st.Request) (st.Request, string) {
	if data.Surl != "" {
		url, found := FindSurl(Catalog, data.Surl)
		if found {
			data.Url = url
			return data, "done"
		} else {
			data.Url = ""
			return data, "url not found"
		}
	} else {
		return data, "requested short_url is empty"
	}
}
