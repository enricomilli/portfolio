package home

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/delaneyj/datastar"
	"github.com/enricomilli/portfolio/lib"
	home_html "github.com/enricomilli/portfolio/site/home/html"
)

type RandomNumberForm struct {
	Low  int `json:"lowRange"`
	High int `json:"highRange"`
}

func HandleRandNum(w http.ResponseWriter, r *http.Request) {

	var formData RandomNumberForm
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		fmt.Println("Decoder Error:", err)
		lib.ResponseWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	fmt.Printf("Form Data: %+v\n", formData)

	sse := datastar.NewSSE(w, r)

	err = datastar.RenderFragmentTempl(sse, home_html.NewRandomNumber(formData.Low+rand.Intn(formData.High-formData.Low), formData.Low, formData.High))
	if err != nil {
		fmt.Println("error rendering fragment:", err)
		lib.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sse.Context().Done()
}

func HandleRenderEditForm(w http.ResponseWriter, r *http.Request) {

	sse := datastar.NewSSE(w, r)

	err := datastar.RenderFragmentTempl(sse, home_html.EditRangeWorkspace())
	if err != nil {
		fmt.Println("error rendering fragment:", err)
		lib.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sse.Context().Done()
}

func HandleUpdatedRange(w http.ResponseWriter, r *http.Request) {

	sse := datastar.NewSSE(w, r)

	err := datastar.RenderFragmentTempl(sse, home_html.GenerateNumsWorkspace())
	if err != nil {
		fmt.Println("error rendering fragment:", err)
		lib.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sse.Context().Done()
}
