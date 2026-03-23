package api

import (
	"UShort/services/url"
	"UShort/types"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store types.Mystore
}

// constructor
func NewAPIServer(addr string, store types.Mystore) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("api/v1").Subrouter()

	urlHandler := url.NewHandler(s.store)
	urlHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
