package main

import (
	"log"
	"net/http"

	"github.com/rusq/aklapi/aklapi"
	"github.com/rusq/osenv/v2"
)

const (
	root           = "/"
	apiRoot        = "/api/v2"
	apiAddress     = apiRoot + "/address/"
	apiDates       = apiRoot + "/dates/"
	apiCollections = apiRoot + "/collections/"
)

var port = osenv.Value("PORT", "8080")

func main() {
	http.HandleFunc(root, aklapi.RootHandler)
	http.HandleFunc(apiAddress, aklapi.AddressHandler)
	http.HandleFunc(apiDates, aklapi.DatesHandler)
	http.HandleFunc(apiCollections, aklapi.CollectionsHandler)

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
