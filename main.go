package main

import (
	"github.com/gin-gonic/gin"
	UserController "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Controllers/User"

)

const (
	port = ":8080"
)

var(
	router = gin.Default()
)

func main(){
	router.GET("/users/:id", UserController.GetUserFromApiC)
	router.GET("/userschannel/:id", UserController.GetUserFromApiChannel)
	router.GET("/usersonechannel/:id", UserController.GetUserFromApiChannelInterface)

	router.Run(port)
}

