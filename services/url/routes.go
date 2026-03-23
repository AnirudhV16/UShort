package url

import (
	"UShort/types"
	"fmt"
	"net/http"

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
	router.HandleFunc("/shorter", h.URLShortenHandler).Methods("Post")
	router.HandleFunc("/:code", h.URLRedirectHandler).Methods("Get")
}

func (h *Handler) URLShortenHandler(w http.ResponseWriter, r *http.Request) {
	bigUrl := r.Body.url
	code := (base64)(int)(h.store["counter"])
	h.store[code] = bigUrl
	h.store["counter"] = (string)((int)(h.store["counter"]) + 1)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(code)
}

func (h *Handler) URLRedirectHandler(w http.ResponseWriter, r *http.Request) error {
	// how to get the query param given in the url???
	shorturl := ""
	v, ok := h.store[shorturl]
	if ok == false {
		return fmt.Errorf("shorturl doesnt exist")
	}
	http.Redirect(w, r, v, 301)
	return fmt.Errorf("")
}
