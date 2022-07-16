package app

import (
	"fmt"
	st "github.com/NautiloosGo/url/internal/storage"
)

// get short url by url requested from catalog
func Get(data st.Request) (st.Request, string) {
	// check for empty request
	if data.Surl != "" {
		// check for symbols from legal list only
		for _, s := range data.Surl {
			if s <= 31 || s >= 127 {
				return data, fmt.Sprintf("Error: wrong symbol %s", string(s))
			}
		}
		// get short url by url requested from catalog
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
