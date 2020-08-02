package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tj/assert"
)

func TestHello(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/raymond", nil)
	w := httptest.NewRecorder()

	GetHello(w, r)

	assert.Equal(t, "Hello, raymond", w.Body.String())
}
