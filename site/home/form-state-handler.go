package home

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/delaneyj/datastar"
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
		fmt.Println("error parsing json", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	sse := datastar.NewSSE(w, r)

	datastar.RenderFragmentTempl(sse, home_html.FormStateServerResponse(formData.FirstName, formData.LastName))

	sse.Context().Done()
}
