package creamcore

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MuxRequestHandler func(http.ResponseWriter, *http.Request)
type RequestHandler func(*http.Request) (int, string)

type Application struct {
	Name   string
	Router *mux.Router
}

func (a *Application) Register(route string, handler RequestHandler, method ...string) *Application {
	a.Router.HandleFunc(route, a.parseResponse(handler)).Methods(method...)
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

func (a *Application) init(name string) {
	a.Name = name
	a.Router = mux.NewRouter()
}

func (a *Application) parseResponse(handler RequestHandler) MuxRequestHandler {
	return func(response http.ResponseWriter, request *http.Request) {
		handlerStatusCode, handlerResponse := handler(request)

		response.WriteHeader(handlerStatusCode)
		response.Write([]byte(handlerResponse))
	}
}
