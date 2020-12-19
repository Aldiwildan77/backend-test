package controllers

import (
	"backend-majoo-test/middlewares"
	"backend-majoo-test/models"
	"backend-majoo-test/services"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	subRouter.Use(middlewares.JwtVerify)
	subRouter.HandleFunc("/", c.createUser).Methods("POST")
	subRouter.HandleFunc("/{id}", c.findUserByID).Methods("GET")
	subRouter.HandleFunc("/", c.findAllUser).Methods("GET")
	subRouter.HandleFunc("/{id}", c.updateUser).Methods("PUT")
	subRouter.HandleFunc("/{id}", c.deleteUser).Methods("DELETE")
	subRouter.HandleFunc("/photo", c.uploadPhoto).Methods("POST")
}

func (c *UserController) createUser(w http.ResponseWriter, r *http.Request) {
	var request models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.UserService.CreateNewUser(request)
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJSON(w, http.StatusCreated, result)
}

func (c *UserController) findUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	authID := r.Context().Value("userToken")
	if int(authID.(float64)) != id {
		ResponseWithError(w, http.StatusForbidden, "forbidden")
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

func (c *UserController) updateUser(w http.ResponseWriter, r *http.Request) {
	var request models.UpdateUserRequest

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	authID := r.Context().Value("userToken")
	if int(authID.(float64)) != id {
		ResponseWithError(w, http.StatusForbidden, "forbidden")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.UserService.UpdateUser(id, request)
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJSON(w, http.StatusOK, result)
}

func (c *UserController) deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	authID := r.Context().Value("userToken")
	if int(authID.(float64)) != id {
		ResponseWithError(w, http.StatusForbidden, "forbidden")
		return
	}

	if err := c.UserService.DeleteUserByID(id); err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJSON(w, http.StatusOK, "User Deleted")
}

func (c *UserController) uploadPhoto(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	authID := r.Context().Value("userToken")
	if int(authID.(float64)) != userID {
		ResponseWithError(w, http.StatusForbidden, "forbidden")
		return
	}

	file, handler, err := r.FormFile("photo")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, _ = io.Copy(f, file)

	path, _ := filepath.Abs(handler.Filename)

	result, err := c.UserService.UploadPhoto(userID, path)
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJSON(w, http.StatusOK, result)
}
