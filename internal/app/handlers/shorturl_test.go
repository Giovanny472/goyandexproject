package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Giovanny472/goyandexproject/internal/app/handlers"
)

func TestShorturl(t *testing.T) {

	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{code: 307, response: "", contentType: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/L3BhcnQ", nil)

			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			h := http.HandlerFunc(handlers.Geturl)
			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()

			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

		})
	}
}
