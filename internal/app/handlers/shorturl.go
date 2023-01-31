package handlers

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
)

var listURL map[string]string

func init() {
	listURL = make(map[string]string)
}

func shortURL(url string) string {

	minlong := 4

	strval := []byte(url)
	newURL := base64.StdEncoding.EncodeToString([]byte(strval))
	if len(newURL) == minlong {
		return newURL
	} else {
		return newURL[:7]
	}
}

func FirstIncrement(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	// GET /{id}
	// принимает  :в качестве URL-параметра идентификатор сокращённого URL
	// возвращает : ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
	case http.MethodGet:
		{
			// принимаем url-параметр
			getShortURL := r.URL.Path
			getShortURL = getShortURL[1:]

			// получаем оригинал-url
			urllong := listURL[getShortURL]
			log.Println("get map val: ", urllong)

			//***** формируем ответ ********
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")

			// код 307
			w.WriteHeader(http.StatusTemporaryRedirect)
			// возвращаем url
			w.Write([]byte(urllong))
		}

	// POST
	// принимает  : в теле запроса строку URL для сокращения
	// возвращает : ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
	case http.MethodPost:
		{
			// получаем url для сокращения
			urllong, er := io.ReadAll(r.Body)
			if er != nil {
				log.Println("post io.readll err: ", er.Error())
			}

			// показать полученный url
			log.Println("long url: ", string(urllong))

			// сокращение url
			shorturl := shortURL(string(urllong))
			listURL[shorturl] = string(urllong)

			log.Println("short url: ", shorturl)

			// код 201
			w.WriteHeader(http.StatusCreated)

			// body
			bd := []byte(shorturl)
			w.Write(bd)
		}

	default:
		w.WriteHeader(http.StatusNotFound)

	}
}
