package url

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/AnirudhV16/UShort/types"
	"github.com/AnirudhV16/UShort/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *Store
}

// constructorrr
func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/shorter", h.URLShortenHandler).Methods("POST")
	router.HandleFunc("/{code}", h.URLRedirectHandler).Methods("GET")
}

func (h *Handler) URLShortenHandler(w http.ResponseWriter, r *http.Request) {
	/*code := h.store["counter"]
	h.store[code] = bigUrl
	temp, _ := strconv.Atoi(h.store["counter"])
	//using string of int numbers as keys
	h.store["counter"] = strconv.Itoa(temp + 1)*/

	//generate a shorturl
	shortUrl := h.generateShorturl()
	//parse the request
	var payload types.URLPayload
	err := utils.ParseJSON(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//check if this short url is in db?
	//if not add this to the db
	// regenerate the short url (5 tries)
	for i := 0; i < 5; i++ {
		err = h.store.AddUrl(types.URL{
			Short_url:    shortUrl,
			Original_url: payload.Original_url,
		})
		if err != nil {
			shortUrl = h.generateShorturl()
		} else {
			break
		}
	}

	//response
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"shorturl": shortUrl})
	/*w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"short_url": "http://localhost:8080/" + shorturl,
	}
	json.NewEncoder(w).Encode(response)*/
}

func (h *Handler) URLRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// how to get the query param given in the url???
	vars := mux.Vars(r)
	shorturl := vars["code"]
	//with short url search db if short url is there get the associated original url and redirect client ot that original url
	origial, err := h.store.GetUrlByShorturl(shorturl)
	if err != nil {
		http.Error(w, "short url not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, origial, http.StatusFound)
}

func (h *Handler) generateShorturl() string {
	shortcode := make([]byte, 6)
	Base62 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	//random 6 characters from "base62"
	//seeding
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		temp := r.Intn(62)
		shortcode[i] = Base62[temp]
	}
	return string(shortcode)
	//code := string(shortcode)
	// not recommended as db query for every shortcode generated can be bottle neck to the database in lage scale
	/*if status, _ := h.store.Checkurl(code); status == false {
		return code
	}*/
}
