package handlers

import (
	"log"
	"net/http"

	"github.com/adam-fraga/adhora/views/layouts"
	"github.com/adam-fraga/adhora/views/metadatas"
	"github.com/adam-fraga/adhora/views/pages"
	"github.com/adam-fraga/adhora/views/partials"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	layout := layouts.Sidebar(
		pages.DashboardPage(),
		metadatas.Head(),
		partials.SideNavbar(),
		partials.Footer(),
	)

	log.Print("DashboardHandler")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	layout.Render(r.Context(), w)
}
