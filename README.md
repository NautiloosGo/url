URL shortener
Create uniq short url. Return original URL by short.

Connection via http. Methods:
POST
waiting:
Content-Type=json or simple form
url:    originalULR

returns:
Content-Type=json
Status 200
Message = errors and comments:
        "done. new short_url"
		"done. already exists"
		"requested url is empty"
Data:  //in case of "done"
    url:        originalULR
    short_url:  newUniqString

GET
waiting:
simple form
short_url:  URLString

returns:
Content-Type=json
Status 200
Message = errors and comments:
		"done"
		"url not found"
		"requested short_url is empty"
Data: //in case of "done"
    url:        originalULR
    short_url:  newUniqString

Parameters can be modified in Confir.json 
    defailt parametres: 
    Short URL length = 10
    Symbols: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_

Database is in memory. Autosave in json file every 10 sec.


