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

func (s APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("api/v1").Subrouter()

	urlStore := url.NewStore(s.db)
	userStore := users.NewStore(s.db)

	urlHandler := url.NewHandler(urlStore)
	urlHandler.RegisterRoutes(subrouter)

	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
