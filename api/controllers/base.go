package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Joel-K-Muraguri/go-jwt/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

type Server struct{
	Router *mux.Router
	DB *gorm.DB
	
}

func (s *Server) Initialize(){


	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)


	s.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", "postgres")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", "postgres")
	}	



	s.DB.Debug().AutoMigrate(&models.User{})

	s.Router = mux.NewRouter()
	
	s.intializeRoutes()

}


func (s *Server) Run(port string){
	fmt.Println("Listening to port 8090")
	log.Fatal(http.ListenAndServe(port, s.Router))
}