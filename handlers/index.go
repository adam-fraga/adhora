package handlers

import (
	t "github.com/a-h/templ"
	layouts "github.com/adam-fraga/adhora/views/layouts"
	pages "github.com/adam-fraga/adhora/views/pages"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	homeData := make(map[string]interface{})
	homeData["Title"] = "Home"

	layout := layouts.Base(pages.Index())

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	t.Handler(layout).ServeHTTP(w, r)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	homeData := make(map[string]interface{})
	homeData["Title"] = "Home"

	layout := layouts.Base(pages.Contact())

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	t.Handler(layout).ServeHTTP(w, r)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	homeData := make(map[string]interface{})
	homeData["Title"] = "Home"

	layout := layouts.Base(pages.About())

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	t.Handler(layout).ServeHTTP(w, r)
}
