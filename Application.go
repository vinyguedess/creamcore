package creamcore

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RequestHandler func(http.ResponseWriter, *http.Request)

type Application struct {
	Name   string
	Router *mux.Router
}

func (a *Application) init(name string) {
	a.Name = name
	a.Router = mux.NewRouter()
}

func (a *Application) Register(method string, route string, handler RequestHandler) *Application {
	a.Router.HandleFunc(route, handler).Methods(method)
	return a
}

func (a *Application) Run(port int) {
	server := http.ListenAndServe(":"+strconv.Itoa(port), a.Router)
	log.Fatal(server)
}

func NewApplication(name string) *Application {
	app := new(Application)
	app.init(name)

	return app
}
