package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

var listURL map[string]string

func shortURL(url string) string {

	strval := []byte(url)
	newURL := base64.StdEncoding.EncodeToString([]byte(strval))
	return newURL
}

func firstIncrement(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	// GET /{id}
	// принимает  :в качестве URL-параметра идентификатор сокращённого URL
	// возвращает : ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
	case http.MethodGet:
		{
			log.Println("get - in...")

			// принимаем url-параметр
			getShortURL := r.URL.Path
			getShortURL = getShortURL[1:]
			fmt.Println("get getShortURL: ", getShortURL)

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
			log.Println("post - in...")

			// получаем url для сокращения
			urllong, erb := io.ReadAll(r.Body)
			if erb != nil {
				log.Println("post io.readll err: ", erb.Error())
			}

			s := "urllongBody: " + string(urllong)
			fmt.Println(s)

			// сокращение url
			newurl := shortURL(string(urllong))

			listURL[newurl] = string(urllong)

			// код 201
			w.WriteHeader(http.StatusCreated)
			//w.Header().Add("content-type", "text/html; charset=UTF-8")

			// body
			bd := []byte(newurl)
			w.Write(bd)
		}

	}
}

func main() {

	listURL = make(map[string]string)

	http.HandleFunc("/", firstIncrement)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
