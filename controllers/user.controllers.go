package controllers

import (
	"backend-majoo-test/services"
	"net/http"

	"github.com/gorilla/mux"
)

// UserController struct
type UserController struct {
	UserService services.UserService
}

// NewUserController to init user controller
func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

// Route entry
func (c *UserController) Route(route *mux.Router) {
	subRouter := route.PathPrefix("/user").Subrouter()

	subRouter.HandleFunc("/", c.findAllUser).Methods("GET")
}

func (c *UserController) findAllUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.UserService.FindAllUser()
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJSON(w, http.StatusOK, result)
}
