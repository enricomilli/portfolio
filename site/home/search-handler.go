package home

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/delaneyj/datastar"
	home_html "github.com/enricomilli/portfolio/site/home/html"
)

type SearchedRequest struct {
	Text string `json:"searchText"`
}

func HandleSearchPage(w http.ResponseWriter, r *http.Request) {

	var searchReq SearchedRequest
	err := json.NewDecoder(r.Body).Decode(&searchReq)
	if err != nil {
	}

	fmt.Printf("Searched for: %s\n", searchReq.Text)

	sse := datastar.NewSSE(w, r)

	searchedFor := strings.TrimSpace(searchReq.Text)

	if searchedFor == "" {
		datastar.RenderFragmentTempl(sse, home_html.HomePage())
	} else {
		datastar.RenderFragmentTempl(sse, home_html.Searched(searchedFor))
	}

	sse.Context().Done()
}
