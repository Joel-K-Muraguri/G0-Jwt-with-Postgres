package api

import (
	"github.com/Joel-K-Muraguri/go-jwt/api/controllers"
)

var server = controllers.Server{}

func Run(){

	server.Initialize()

	server.Run(":8090")

}