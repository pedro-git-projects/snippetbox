package main

import (
	"net/http"
	"testing"

	"github.com/pedro-git-projects/snippetbox/internal/models/assert"
)

func TestPing(t *testing.T) {

	app := newTestApplication(t)

	testServer := newTestServer(t, app.routes())
	defer testServer.Close()

	code, _, body := testServer.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
