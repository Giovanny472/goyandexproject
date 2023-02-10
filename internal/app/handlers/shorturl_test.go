package handlers_test

import (
	"net/http"
	"testing"
)

func TestShorturlGet(t *testing.T) {

	url := "http://127.0.0.1:8080/L3BhcnQ"
	req, e := http.Get(url)
	if e != nil {
		t.Error("error get : ", e.Error())
	}
	defer req.Body.Close()

	if req.StatusCode != 307 {
		t.Errorf(" expecte code 307; got %d", req.StatusCode)
	}

}

/*
func TestShorturlPost(t *testing.T) {

	url := "http://127.0.0.1:8080"
	param := strings.NewReader("/part01/section01/users01")

	req, er := http.NewRequest(http.MethodPost, url, param)
	if er != nil {
		t.Error("error NewRequest-Post: ", er.Error())
	}

	req.Header.Add("content-type", "text/html; charset=UTF-8")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		t.Errorf(" expected code 201; got %d", res.StatusCode)
	}

}
*/
