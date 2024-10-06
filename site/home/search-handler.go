package home

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/delaneyj/datastar"
	"github.com/enricomilli/portfolio/lib"
	home_html "github.com/enricomilli/portfolio/site/home/html"
)

type SearchedRequest struct {
	Text string `json:"searchText"`
}

func HandleSearchPage(w http.ResponseWriter, r *http.Request) {

	var searchReq SearchedRequest
	err := json.NewDecoder(r.Body).Decode(&searchReq)
	if err != nil {
		lib.ResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("could not decode request body: %v", err))
		return
	}

	fmt.Printf("Searched for: %s\n", searchReq.Text)

	sse := datastar.NewSSE(w, r)

	searchedFor := strings.TrimSpace(searchReq.Text)

	if searchedFor == "" {
		err = datastar.RenderFragmentTempl(sse, home_html.HomePage())
		if err != nil {
			lib.ResponseWithError(w, http.StatusInternalServerError, "error sending hypermedia")
			return
		}
	} else {
		err = datastar.RenderFragmentTempl(sse, home_html.SearchPage(searchedFor))
		if err != nil {
			lib.ResponseWithError(w, http.StatusInternalServerError, "error sending hypermedia")
			return
		}
	}

	sse.Context().Done()
}
