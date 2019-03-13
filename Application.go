package creamcore

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// MuxRequestHandler defines Mux Request Handler typo
type MuxRequestHandler func(http.ResponseWriter, *http.Request)

// RequestHandler defines CReaM Request Handler typo
type RequestHandler func(*http.Request) (int, string)

// NewApplication is a Singleton responsible of returning Application instance
func NewApplication(name string) *Application {
	app := new(Application)
	app.init(name)

	return app
}

// Application is the core class responsible for registering handlers,
// middlewares and start up the app
type Application struct {
	Name   string
	Router *mux.Router
}

// Register is the method responsible of registering middlewares
func (a *Application) Register(route string, handler RequestHandler, method ...string) *Application {
	a.Router.HandleFunc(route, a.parseResponse(handler)).Methods(method...)
	return a
}

// Run is the method responsible of  starting up the application server
func (a *Application) Run(port int) {
	server := http.ListenAndServe(":"+strconv.Itoa(port), a.Router)
	log.Fatal(server)
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
