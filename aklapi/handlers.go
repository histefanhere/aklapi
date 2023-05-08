package aklapi

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
)

const dateFmt = "2006-01-02"

var tmpl = template.Must(template.New("index.html").Parse(rootHTML))

type datesResponse struct {
	Rubbish    string `json:"rubbish,omitempty"`
	Recycle    string `json:"recycle,omitempty"`
	FoodScraps string `json:"foodscraps,omitempty"`
	Address    string `json:"address,omitempty"`
}

type errorResponse struct {
	Error string `json:"error,omitempty"`
}

func respond(w http.ResponseWriter, data interface{}, code int) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(b)
}

func rubbish(r *http.Request) (*CollectionDayDetailResult, error) {
	addr := r.FormValue("addr")
	if addr == "" {
		return nil, errors.New(http.StatusText(http.StatusBadRequest))
	}
	return CollectionDayDetail(addr)
}

// Handles the endpoint for looking up the address key and full address for a partial address
func AddressHandler(w http.ResponseWriter, r *http.Request) {
	addr := r.FormValue("addr")
	if addr == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	resp, err := AddressLookup(addr)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	respond(w, resp, http.StatusOK)
}

// Handles the endpoint for looking up the next collection dates for each type of waste
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	res, err := rubbish(r)
	if err != nil {
		respond(w, &errorResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	resp := datesResponse{
		Recycle:    res.NextRecycle().Format(dateFmt),
		Rubbish:    res.NextRubbish().Format(dateFmt),
		FoodScraps: res.NextFoodScraps().Format(dateFmt),
		Address:    res.Address.Address,
	}
	respond(w, resp, http.StatusOK)
}

// Handles the endpoint for more information about each collection and the address
func CollectionsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := rubbish(r)
	if err != nil {
		respond(w, errorResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	respond(w, res, http.StatusOK)
}

// Handles the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	var page = struct {
		CyberdyneLogo string
	}{
		CyberdyneLogo: cyberdynePng,
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", &page); err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
}
