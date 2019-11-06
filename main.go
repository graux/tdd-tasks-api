package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {
	startTasksApiServer(8080)
}

func startTasksApiServer(serverPort int) *ApiServer {
	router := NewRouter()
	// Añadir rutas aquí
	// router.AddRoute("/task", getTasks, http.MethodGet)
	router.ServeStaticFolder("static")
	apiServer := NewApiServer(serverPort, router)
	apiServer.Start()
	return apiServer
}

type ApiServer struct {
	server *http.Server
	router *Router
}

type Router struct {
	handler *mux.Router
}

func NewApiServer(serverPort int, router *Router) *ApiServer {
	apiServer := &ApiServer{
		server: &http.Server{Addr: fmt.Sprintf(":%d", serverPort)},
		router: router,
	}
	return apiServer
}

func (server *ApiServer) Start() {
	go func() {
		if err := server.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Listen and Server Error: %s", err)
		}
	}()
}

func (server *ApiServer) Stop() {
	if err := server.server.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}

func NewRouter() *Router {
	muxHandler := mux.NewRouter().StrictSlash(true)
	router := &Router{
		handler: muxHandler,
	}
	return router
}

func (rt *Router) AddRoute(ruta string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	rt.handler.HandleFunc(ruta, handler).Methods(methods...)
}

func (rt *Router) HandleRequest(response http.ResponseWriter, request *http.Request) {
	rt.ServeHTTP(response, request)
}

func (rt *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	rt.handler.ServeHTTP(response, request)
}

func (rt *Router) ServeStaticFolder(staticRelFolder string) {
	basePath, _ := filepath.Abs("./")
	staticPath := path.Join(basePath, staticRelFolder)
	fileServer := http.FileServer(http.Dir(staticPath))
	rt.handler.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))
}
