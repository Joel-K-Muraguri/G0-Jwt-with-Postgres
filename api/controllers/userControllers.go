package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Joel-K-Muraguri/go-jwt/api/models"
	"github.com/Joel-K-Muraguri/go-jwt/api/responses"
	"github.com/gorilla/mux"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "This is Go JWT with Postgres API")

}


func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(s.DB)

	if err != nil {

		formattedError := responses.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
	// responses.JSON(w, http.StatusOK, "create a user")

}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateUser(s.DB, uint32(uid))
	if err != nil {
		formattedError := responses.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
	// responses.JSON(w, http.StatusOK, "Update a user")

}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = user.DeleteUser(s.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
	// responses.JSON(w, http.StatusOK, "Delete a user")

}

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request){
	user := models.User{}

	users, err := user.FindAllUsers(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
	// responses.JSON(w, http.StatusOK, "Get all users")

}

func (s *Server) GetAUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	users := models.User{}
	userGotten, err := users.FindAUser(s.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
	// responses.JSON(w, http.StatusOK, "Get a user")

}