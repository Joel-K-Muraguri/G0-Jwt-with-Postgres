package api

import (
	"github.com/Joel-K-Muraguri/go-jwt/api/controllers"
	"github.com/Joel-K-Muraguri/go-jwt/api/seed"
)

var server = controllers.Server{}

func Run(){

	server.Initialize()

	seed.Load(server.DB)

	server.Run(":8090")

}