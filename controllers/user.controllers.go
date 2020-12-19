package controllers

import (
	"backend-majoo-test/services"
	"net/http"
	"strconv"

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

	subRouter.HandleFunc("/{id}", c.findUserByID).Methods("GET")
	subRouter.HandleFunc("/", c.findAllUser).Methods("GET")
}

func (c *UserController) findUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.UserService.FindUserByID(id)
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJSON(w, http.StatusOK, result)
}

func (c *UserController) findAllUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.UserService.FindAllUser()
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJSON(w, http.StatusOK, result)
}
