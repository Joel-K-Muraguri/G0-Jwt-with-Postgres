package controllers

import(
	"net/http"
	"github.com/Joel-K-Muraguri/go-jwt/api/responses"

)

func (s *Server) Home(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "This is Go JWT with Postgres API")

}


func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "create a user")

}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "Update a user")

}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "Delete a user")

}

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "Get all users")

}

func (s *Server) GetAUser(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "Get a user")

}