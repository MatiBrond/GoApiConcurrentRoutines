package User

import (
	"github.com/gin-gonic/gin"
	Explosure2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Domain/Explosure"
	Site2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Domain/Site"
	User2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Domain/User"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Services/Explosure"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Services/Site"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Services/User"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Utils/ApiErrors"
	"net/http"
	"strconv"
	"sync"
)

const(

	paramUserID = "id"
)

type UserSite struct{
	User *User2.User
	Site *Site2.Site
	Explosure *Explosure2.Explosure_level
}

func GetUserFromApiC(context *gin.Context) {

	var result UserSite
	var wg sync.WaitGroup

	id := context.Param(paramUserID)
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &ApiErrors.ApiError{
			Message: "Fatal URL",
			Status:  http.StatusBadRequest}
		context.JSON(apiError.Status, apiError)
		return
	}

	user, apiError := User.GetUserFromApi(userID);
	if apiError != nil {
		context.JSON(apiError.Status, apiError)
		return
	}
	result.User = user

	wg.Add(2)
	go func(){
		defer wg.Done()
		site, err1 := Site.GetSiteFromApi(user.SiteID)
		if err1 != nil {
			apiError := &ApiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		result.Site = site
	}()

	go func () {
		defer wg.Done()
		explosure, err2 := Explosure.GetUserFromApi(user.SiteID)
		if err2 != nil {
			apiError := &ApiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		result.Explosure = explosure
	}()

	wg.Wait()
	context.JSON(200, &result)
}

func GetUserFromApiChannel(context *gin.Context) {

	var result UserSite
	siteChan := make(chan *Site2.Site)
	explosureChan := make(chan *Explosure2.Explosure_level)

	id := context.Param(paramUserID)
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &ApiErrors.ApiError{
			Message: "Fatal URL",
			Status:  http.StatusBadRequest}
		context.JSON(apiError.Status, apiError)
		return
	}

	user, apiError := User.GetUserFromApi(userID);
	if apiError != nil {
		context.JSON(apiError.Status, apiError)
		return
	}
	result.User = user

	go func(){
		site, err1 := Site.GetSiteFromApi(user.SiteID)
		if err1 != nil {
			apiError := &ApiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		siteChan <- site
	}()

	go func () {
		explosure, err2 := Explosure.GetUserFromApi(user.SiteID)
		if err2 != nil {
			apiError := &ApiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		explosureChan <- explosure
	}()

	result.Site = <- siteChan
	result.Explosure = <- explosureChan
	context.JSON(200, &result)
}



