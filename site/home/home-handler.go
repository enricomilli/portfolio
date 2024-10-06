package home

import (
	"fmt"
	"net/http"

	"github.com/delaneyj/datastar"
	"github.com/enricomilli/portfolio/lib"
	home_html "github.com/enricomilli/portfolio/site/home/html"
)

func HandleHomePage(w http.ResponseWriter, r *http.Request) {

	sse := datastar.NewSSE(w, r)

	err := datastar.RenderFragmentTempl(sse, home_html.HomePage())
	if err != nil {
		fmt.Println(err)
		lib.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	sse.Context().Done()
}
