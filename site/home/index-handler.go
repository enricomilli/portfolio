package home

import (
	"net/http"

	home_html "github.com/enricomilli/portfolio/site/home/html"
)

// Handles initial page load
func HandleIndex(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	home_html.Page().Render(ctx, w)

}
