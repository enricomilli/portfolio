package home

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/delaneyj/datastar"
	"github.com/enricomilli/portfolio/lib"
	home_html "github.com/enricomilli/portfolio/site/home/html"
)

type FormStateData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func HandleFormStatePut(w http.ResponseWriter, r *http.Request) {

	var formData FormStateData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		lib.ResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
	}

	sse := datastar.NewSSE(w, r)

	datastar.RenderFragmentTempl(sse, home_html.FormStateServerResponse(formData.FirstName, formData.LastName))
	if err != nil {
		lib.ResponseWithError(w, http.StatusInternalServerError, fmt.Sprintf("error rendering hypermedia: %v", err))
	}

	sse.Context().Done()
}
