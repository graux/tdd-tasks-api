package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_startTasksApiServer(t *testing.T) {
	apiServer := startTasksApiServer(8083)
	apiServer.Stop()
}

func Test_main_struct(t *testing.T) {
	apiServer := NewApiServer(8081, nil)
	if apiServer == nil {
		t.Error("apiServer no tiene que ser null")
	}
}

func Test_main_crear_router(t *testing.T) {

	router := NewRouter()
	if router == nil {
		t.Error("router no tiene que ser null")
	}
}
func Test_main_asignar_router(t *testing.T) {
	router := NewRouter()
	apiServer := NewApiServer(8081, router)
	if apiServer.router == nil {
		t.Error("Server should have a router")
	}
}

func Test_main_test_add_route(t *testing.T) {
	router := NewRouter()
	router.AddRoute("prueba", nil)
}

func Test_main_apiserver_test_route(t *testing.T) {
	router := NewRouter()
	router.AddRoute("/prueba", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}, http.MethodGet)

	request := httptest.NewRequest(http.MethodGet, "/prueba", nil)
	response := httptest.NewRecorder()

	router.handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Status code should be %d, got %d", http.StatusOK, response.Code)
	}
}

func Test_routerHandleRequest(t *testing.T) {
	router := NewRouter()
	router.AddRoute("/prueba", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}, http.MethodGet)

	request := httptest.NewRequest(http.MethodGet, "/prueba", nil)
	response := httptest.NewRecorder()

	router.HandleRequest(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Status code should be %d, got %d", http.StatusOK, response.Code)
	}
}

func Test_staticRouteNotFound(t *testing.T) {
	router := NewRouter()
	router.ServeStaticFolder("static")
	request := httptest.NewRequest(http.MethodGet, "/notfound.html", nil)
	response := httptest.NewRecorder()

	router.HandleRequest(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Status code should be %d, got %d", http.StatusNotFound, response.Code)
	}
}

func Test_staticRouteOK(t *testing.T) {
	router := NewRouter()
	router.ServeStaticFolder("static")
	request := httptest.NewRequest(http.MethodGet, "/favicon.ico", nil)
	response := httptest.NewRecorder()

	router.HandleRequest(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Status code should be %d, got %d: %v", http.StatusOK, response.Code, response.HeaderMap)
	}
}
