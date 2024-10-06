package home

import (
	"fmt"
	"net/http"

	"github.com/enricomilli/portfolio/lib"
	home_html "github.com/enricomilli/portfolio/site/home/html"
)

// Handles initial page load
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := home_html.Page().Render(ctx, w)
	if err != nil {
		lib.ResponseWithError(w, http.StatusInternalServerError, fmt.Sprintf("error rendering hypermedia: %v", err))
	}
}
