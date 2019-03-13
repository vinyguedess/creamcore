package creamcore

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// MuxRequestHandler defines Mux Request Handler typo
type MuxRequestHandler func(http.ResponseWriter, *http.Request)

// RequestHandler defines CReaM Request Handler typo
type RequestHandler func(Request) (int, string)

// Request is the HTTP request for handler
type Request struct {
	URL    string
	Method string
	Header http.Header
	Body   io.ReadCloser
}

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
	Server *http.Server
}

// Register is the method responsible of registering middlewares
func (a *Application) Register(route string, handler RequestHandler, method ...string) *Application {
	a.Router.HandleFunc(route, a.parseResponse(handler)).Methods(method...)
	return a
}

// Run is the method responsible of  starting up the application server
func (a *Application) Run(port int) {
	a.Server = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: a.Router,
	}
	a.Server.ListenAndServe()
	log.Fatal(a.Server)
}

func (a *Application) init(name string) {
	a.Name = name
	a.Router = mux.NewRouter()
}

func (a *Application) parseRequest(request *http.Request) Request {
	return Request{
		URL:    fmt.Sprintf("%s/%s", request.URL.Host, request.URL.Path),
		Method: request.Method,
		Header: request.Header,
		Body:   request.Body,
	}
}

func (a *Application) parseResponse(handler RequestHandler) MuxRequestHandler {
	return func(response http.ResponseWriter, request *http.Request) {
		handlerStatusCode, handlerResponse := handler(a.parseRequest(request))

		response.WriteHeader(handlerStatusCode)
		response.Write([]byte(handlerResponse))
	}
}
