package api

import (
	"database/sql"
	"net/http"

	"github.com/AnirudhV16/UShort/services/url"
	"github.com/AnirudhV16/UShort/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	//store types.Mystore
	db *sql.DB
}

// constructor
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	println(1)
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	println(2)
	urlStore := url.NewStore(s.db)
	userStore := users.NewStore(s.db)
	println(3)
	urlHandler := url.NewHandler(urlStore)
	urlHandler.RegisterRoutes(subrouter)
	println(4)
	userHandler := users.NewHandler(userStore)
	println(5)
	userHandler.RegisterRoutes(subrouter)
	println(6)
	println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
