package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedro-git-projects/snippetbox/internal/models/assert"
)

func TestPing(t *testing.T) {

	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}

	testServer := httptest.NewTLSServer(app.routes())
	defer testServer.Close()

	result, err := testServer.Client().Get(testServer.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, result.StatusCode, http.StatusOK)

	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
