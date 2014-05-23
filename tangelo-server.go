package main

import (
	"encoding/json"
	"fmt"
	"github.com/rdxiang/tangelo-server/router"
	"net/http"
)

type APIVersion struct {
	VersionString string `json:"version"`
}

const APIVersionString = "v0.0.1"

//Make sure to use http constants

func getAPIInformation(w http.ResponseWriter, r *http.Request) {
	var response APIResponse
	var apiVersion APIVersion
	var err error
	apiVersion.VersionString = APIVersionString
	response.Data, err = json.Marshal(apiVersion)
	if err != nil {
		fmt.Println("Error:", err)
	}
	response.send(w)
}

func sendEmpty(w http.ResponseWriter, r *http.Request) {
	var response APIResponse
	response.send(w)

}

func setupRoutes() {
	router.AddAPIHandler("/", getAPIInformation)
	router.AddAPIHandler("/empty", sendEmpty)
}

func main() {
	setupRoutes()
	setupTangeloRoutes()
	http.Handle("/", router.Router)
	http.ListenAndServe("localhost:4000", nil)
}
