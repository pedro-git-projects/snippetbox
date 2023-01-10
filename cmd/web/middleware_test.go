package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedro-git-projects/snippetbox/internal/models/assert"
)

func TestSecureHeaders(t *testing.T) {
	recorder := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(recorder, r)
	result := recorder.Result()

	expected := "origin-when-cross-origin"
	assert.Equal(t, result.Header.Get("Referrer-Policy"), expected)

	expected = "nosniff"
	assert.Equal(t, result.Header.Get("X-Content-Type-Options"), expected)

	expected = "deny"
	assert.Equal(t, result.Header.Get("X-Frame-Options"), expected)

	expected = "0"
	assert.Equal(t, result.Header.Get("X-XSS-Protection"), expected)

	assert.Equal(t, result.StatusCode, http.StatusOK)

	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")

}
