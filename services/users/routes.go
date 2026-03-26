package users

import (
	"fmt"
	"net/http"

	"github.com/AnirudhV16/UShort/config"
	"github.com/AnirudhV16/UShort/services/auth"
	"github.com/AnirudhV16/UShort/types"
	"github.com/AnirudhV16/UShort/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *Store
}

// constructorr
func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/Login", h.LoginHandler).Methods("POST")
	router.HandleFunc("/Register", h.RegisterHandler).Methods("POST")
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	err := utils.ParseJSON(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//get the user frome email
	user, _ := h.store.GetUserByGmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user credentials wrong!!", payload.Email))
	}
	//compare passwords from db if they match login else error

	if !(auth.Compare(user.Password, []byte(payload.Password))) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user credentials wrong!!", payload.Email))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var payload types.RegisterUserPayload
	err := utils.ParseJSON(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//check if user exists??
	_, err = h.store.GetUserByGmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	//creating  a new user if the user doesnt exist in the user store
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
