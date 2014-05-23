package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router
var APIRouter *mux.Router

type ResponseFunction func(w http.ResponseWriter, r *http.Request)

func init() {
	makeRouter()
}

func makeRouter() {
	Router = mux.NewRouter()
	makeAPIRouter()
}

func AddAPIHandler(path string, handler ResponseFunction) {
	APIRouter.HandleFunc(path, handler)
}

func makeAPIRouter() {
	APIRouter = Router.PathPrefix("/api").Subrouter()

}
