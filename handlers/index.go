package handlers

import (
	"log"
	"net/http"

	loc "github.com/adam-fraga/adhora/localized"
	layouts "github.com/adam-fraga/adhora/views/layouts"
	"github.com/adam-fraga/adhora/views/metadatas"
	"github.com/adam-fraga/adhora/views/pages"
	"github.com/adam-fraga/adhora/views/partials"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var homeData = loc.GetHomePageDataEn()

	navlinks, ok := homeData["nav_links"].([]string)
	if !ok {
		log.Fatal("Error getting nav links")
	}

	layout := layouts.Base(
		pages.Index(),
		metadatas.Head(),
		partials.Header(navlinks),
		partials.Footer(),
	)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	layout.Render(r.Context(), w)
}
