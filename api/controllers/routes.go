package controllers

func (s *Server) intializeRoutes(){

	s.Router.HandleFunc("users/home", s.Home).Methods("GET")
	s.Router.HandleFunc("users/create", s.CreateUser).Methods("POST")
	s.Router.HandleFunc("users/all", s.GetAllUsers).Methods("GET")
	s.Router.HandleFunc("users/{id}", s.GetAUser).Methods("GET")
	s.Router.HandleFunc("users/update/{id}", s.UpdateUser).Methods("PUT")
	s.Router.HandleFunc("users/delete./{id}", s.DeleteUser).Methods("DELETE")

}