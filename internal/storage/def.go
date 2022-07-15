package storage

type Request struct {
	Id   string `json:"id"`
	Url  string `json:"url"`
	Surl string `json:"surl"`
}

type Settings struct {
	Letters string
	Qty     int
}

type Config struct {
	Settings struct {
		Letters string `json:"letters"`
		Qty     int    `json:"url_len"`
	} `json:"settings"`
	FileCatalog string `json:"local_database"`
}

type Catalog struct {
	List []Request `json:"links"`
}
