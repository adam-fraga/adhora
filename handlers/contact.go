package handlers

import (
	"log"
	"net/http"

	t "github.com/a-h/templ"
	loc "github.com/adam-fraga/adhora/localized"
	layouts "github.com/adam-fraga/adhora/views/layouts"
	"github.com/adam-fraga/adhora/views/metadatas"
	"github.com/adam-fraga/adhora/views/pages"
	"github.com/adam-fraga/adhora/views/partials"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	homeData := loc.GetHomePageDataEn()

	navlinks, ok := homeData["nav_links"].([]string)
	if !ok {
		log.Fatal("Error getting nav links")
	}

	layout := layouts.Base(
		pages.Contact(),
		metadatas.Head(),
		partials.Header(navlinks),
		partials.Footer(),
	)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	t.Handler(layout).ServeHTTP(w, r)
}
