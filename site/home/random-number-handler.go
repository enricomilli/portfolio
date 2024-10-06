package home

import (
	"math/rand"
	"net/http"

	"github.com/delaneyj/datastar"
	"github.com/enricomilli/portfolio/lib"
	home_html "github.com/enricomilli/portfolio/site/home/html"
)

func HandleRandNum(w http.ResponseWriter, r *http.Request) {

	sse := datastar.NewSSE(w, r)

	err := datastar.RenderFragmentTempl(sse, home_html.RandomNumber(rand.Intn(100)))
	if err != nil {
		lib.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	sse.Context().Done()
}
