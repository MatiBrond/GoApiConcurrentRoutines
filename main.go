package main

import (
	"github.com/gin-gonic/gin"
	SiteController "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/controllers/site"
	UserController "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/controllers/user"
)

const (
	port = ":8083"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/users/:id", UserController.GetUserFromApiC)
	router.GET("/userschannel/:id", UserController.GetUserFromApiChannel)
	router.GET("/usersonechannel/:id", UserController.GetUserFromApiChannelInterface)
	router.GET("/usersmock/", UserController.GetMock)
	//router.GET("/siteCB/", SiteController.GetSiteFromApiCB)
	router.GET("/siteDBF", SiteController.GetSiteFromApiCB)

	router.Run(port)
}
