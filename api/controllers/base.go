package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct{
	Router *mux.Router
	
}

func (s *Server) Initialize(){

}


func (s *Server) Run(port string){
	fmt.Println("Listening to port 8090")
	log.Fatal(http.ListenAndServe(port, s.Router))
}