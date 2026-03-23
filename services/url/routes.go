package url

import (
	"UShort/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.Mystore
}

// constructorrr
func NewHandler(store types.Mystore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/shorter", h.URLShortenHandler).Methods("POST")
	router.HandleFunc("/{code}", h.URLRedirectHandler).Methods("GET")
}

func (h *Handler) URLShortenHandler(w http.ResponseWriter, r *http.Request) {
	var req types.Requestt
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	bigUrl := req.Url
	code := h.store["counter"]
	h.store[code] = bigUrl
	temp, _ := strconv.Atoi(h.store["counter"])
	//using string of int numbers as keys
	h.store["counter"] = strconv.Itoa(temp + 1)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"short_url": "http://localhost:8080/" + code,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) URLRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// how to get the query param given in the url???
	vars := mux.Vars(r)
	shorturl := vars["code"]
	v, ok := h.store[shorturl]
	if ok == false {
		http.Error(w, "short url not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, v, http.StatusFound)
}
