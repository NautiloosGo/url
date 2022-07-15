package app

import (
	st "github.com/NautiloosGo/url/internal/storage"
)

func Post(data st.Request) (st.Request, string) {
	if data.Url != "" {
		surl, found := FindUrl(Catalog, data.Url)
		if found {
			data.Surl = surl
			return data, "done. already exists"
		} else {
			return PostUniq(data)
		}
	} else {
		return data, "requested url is empty"
	}
}

func PostUniq(data st.Request) (st.Request, string) {
	surl := GetRandomString(Conf.Settings.Qty, Conf.Settings.Letters)
	if _, found := FindSurl(Catalog, surl); found {
		return PostUniq(data)
	} else {
		data.Surl = surl
		AddLink(data)
		return data, "done. new short_url"
	}
}

func AddLink(data st.Request) {
	Catalog.List = append(Catalog.List, data)
}
