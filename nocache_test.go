package wnocache

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_NocacheHeaders(t *testing.T) {
	responseHeaders := map[string]string{
		"Cache-Control": "no-cache, no-store, max-age=0, must-revalidate, no-transform",
		"Pragma":        "no-cache",
		"Expires":       expires,
	}
	w := httptest.NewRecorder()
	g := gin.New()

	g.Use(New())
	r, _ := http.NewRequest("GET", "/", nil)
	g.ServeHTTP(w, r)

	for key, value := range responseHeaders {
		if w.Header()[key][0] != value {
			t.Errorf("Missing header: %s", key)
		}
	}
}
