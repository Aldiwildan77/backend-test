package controllers

import (
	"backend-majoo-test/models"
	"backend-majoo-test/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AuthController struct
type AuthController struct {
	AuthService services.AuthService
}

// NewAuthController to init auth controller
func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{authService}
}

// Route entry
func (c *AuthController) Route(route *mux.Router) {
	subRouter := route.PathPrefix("/login").Subrouter()

	subRouter.HandleFunc("/", c.login).Methods("POST")
}

func (c *AuthController) login(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.AuthService.Login(request)
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJSON(w, http.StatusOK, result)
}
