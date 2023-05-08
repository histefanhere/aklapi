package main

import (
	"log"
	"net/http"

	"github.com/rusq/aklapi/aklapi"
	"github.com/rusq/osenv/v2"
)

const (
	root              = "/"
	apiRoot           = "/api/v1"
	apiAddr           = apiRoot + "/addr/"
	apiRubbishRecycle = apiRoot + "/rr/"
	apiRRExt          = apiRoot + "/rrext/"
)

var port = osenv.Value("PORT", "8080")

func main() {
	http.HandleFunc(root, aklapi.RootHandler)
	http.HandleFunc(apiAddr, aklapi.AddrHandler)
	http.HandleFunc(apiRubbishRecycle, aklapi.RrHandler)
	http.HandleFunc(apiRRExt, aklapi.RrExtHandler)

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
