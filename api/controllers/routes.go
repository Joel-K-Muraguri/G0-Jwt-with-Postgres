package controllers

import (
	"github.com/Joel-K-Muraguri/go-jwt/api/middleware"
)

func (s *Server) intializeRoutes(){

	s.Router.HandleFunc("/users/home", middleware.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/users/create", middleware.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/all", middleware.SetMiddlewareJSON(s.GetAllUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJSON(s.GetAUser)).Methods("GET")
	s.Router.HandleFunc("/users/update/{id}", middleware.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/users/delete/{id}", middleware.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")

}