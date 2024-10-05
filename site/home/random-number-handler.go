package home

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/delaneyj/datastar"
	home_html "github.com/enricomilli/portfolio/site/home/html"
)

func HandleRandNum(w http.ResponseWriter, r *http.Request) {

	sse := datastar.NewSSE(w, r)

	err := datastar.RenderFragmentTempl(sse, home_html.RandomNumber(rand.Intn(100)))
	if err != nil {
		fmt.Println(err)
	}

	sse.Context().Done()
}
