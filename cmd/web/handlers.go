package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pedro-git-projects/snippetbox/internal/models"
)

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	templateSet, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buffer := new(bytes.Buffer)

	err := templateSet.ExecuteTemplate(buffer, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

	w.WriteHeader(status)
	buffer.WriteTo(w)

}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.tmpl", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form to create new snippet"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "Who’s joe?"
	content := `"Who’s joe?" a distant voice asks.\n a horrific creature crawled from the ground, covered in mucus and tar.\n ”Joe Momma…” the creature whispered.`
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
