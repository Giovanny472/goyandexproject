package handlers

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

const (
	prefixURL string = "http://gio.com/"
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

func RouterInc(r chi.Router) {

	r.Get("/{id}", geturl)

	r.Post("/", posturl)

}

func geturl(w http.ResponseWriter, r *http.Request) {

	// принимаем url-параметр
	strShortURL := chi.URLParam(r, "id")
	sliceShortURL := strings.Split(strShortURL, " ")
	getShortURL := sliceShortURL[1]

	// получаем оригинал-url
	urllong := listURL[getShortURL]
	log.Print("[GET] getUrlong:", urllong, " ,shortURL:", getShortURL)

	//***** формируем ответ ********
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	// код 307/ 301
	w.WriteHeader(http.StatusMovedPermanently)
	// возвращаем url
	w.Write([]byte(urllong))
}

func posturl(w http.ResponseWriter, r *http.Request) {

	// получаем url для сокращения
	urllong, er := io.ReadAll(r.Body)
	if er != nil {
		log.Println("post io.readll err body: ", er.Error())
	}

	// показать полученный url

	// сокращение url
	encod := shortURL(string(urllong))
	shorturl := prefixURL + encod
	listURL[encod] = string(urllong)
	log.Print("[POST] short-url:", shorturl, " long-url:", string(urllong), ", encod:", encod)

	// код 201
	w.WriteHeader(http.StatusCreated)

	// body
	bd := []byte(shorturl)
	w.Write(bd)
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

			// код 307 / 301
			w.WriteHeader(http.StatusMovedPermanently)
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
