package user

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	Explosure2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/domain/explosure"
	Site2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/domain/site"
	User2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/domain/user"
	Explosure "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/services/explosure"
	Site "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/services/site"
	User "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/services/user"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/utils/apiErrors"
)

const (
	paramUserID = "id"
)

type UserSite struct {
	User      *User2.User
	Site      *Site2.Site
	Explosure *Explosure2.Explosure_level
}

func GetUserFromApiC(context *gin.Context) {

	var result UserSite
	var wg sync.WaitGroup

	id := context.Param(paramUserID)
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &apiErrors.ApiError{
			Message: "Fatal URL",
			Status:  http.StatusBadRequest}
		context.JSON(apiError.Status, apiError)
		return
	}

	user, apiError := User.GetUserFromApi(userID)
	if apiError != nil {
		context.JSON(apiError.Status, apiError)
		return
	}
	result.User = user

	wg.Add(2)
	go func() {
		defer wg.Done()
		site, err1 := Site.GetSiteFromApi(user.SiteID)
		if err1 != nil {
			apiError := &apiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		result.Site = site
	}()

	go func() {
		defer wg.Done()
		explosure, err2 := Explosure.GetExpFromApi(user.SiteID)
		if err2 != nil {
			apiError := &apiErrors.ApiError{
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
		apiError := &apiErrors.ApiError{
			Message: "Fatal URL",
			Status:  http.StatusBadRequest}
		context.JSON(apiError.Status, apiError)
		return
	}

	user, apiError := User.GetUserFromApi(userID)
	if apiError != nil {
		context.JSON(apiError.Status, apiError)
		return
	}
	result.User = user

	go func() {
		site, err1 := Site.GetSiteFromApi(user.SiteID)
		if err1 != nil {
			apiError := &apiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		siteChan <- site
	}()

	go func() {
		explosure, err2 := Explosure.GetExpFromApi(user.SiteID)
		if err2 != nil {
			apiError := &apiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		explosureChan <- explosure
	}()

	result.Site = <-siteChan
	result.Explosure = <-explosureChan
	context.JSON(200, &result)
}

func GetUserFromApiChannelInterface(context *gin.Context) {

	var result UserSite
	var wg sync.WaitGroup
	c := make(chan UserSite, 3)
	var r UserSite

	id := context.Param(paramUserID)
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &apiErrors.ApiError{
			Message: "Fatal URL",
			Status:  http.StatusBadRequest}
		context.JSON(apiError.Status, apiError)
		return
	}

	user, apiError := User.GetUserFromApi(userID)
	if apiError != nil {
		context.JSON(apiError.Status, apiError)
		return
	}
	r.User = user
	c <- r

	wg.Add(2)
	go func() {
		defer wg.Done()
		var r UserSite
		site, err1 := Site.GetSiteFromApi(user.SiteID)
		if err1 != nil {
			apiError := &apiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		r.Site = site
		c <- r
	}()

	go func() {
		defer wg.Done()
		var r UserSite
		explosure, err2 := Explosure.GetExpFromApi(user.SiteID)
		if err2 != nil {
			apiError := &apiErrors.ApiError{
				Message: "Fatal URL",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		r.Explosure = explosure
		c <- r
	}()
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:

			if r.Site != nil {
				result.Site = r.Site
				continue
			}
			if r.Explosure != nil {
				result.Explosure = r.Explosure
				continue
			}
			if r.User != nil {
				result.User = r.User
				continue
			}
		}
	}

	wg.Wait()
	context.JSON(200, &result)
}

func GetMock(context *gin.Context) {

	var result UserSite
	var wg sync.WaitGroup
	c := make(chan UserSite, 3)
	var r UserSite

	user, apiError := User.GetUserFromMock()
	if apiError != nil {
		context.JSON(apiError.Status, apiError)
		return
	}
	r.User = user
	c <- r

	wg.Add(2)
	go func() {
		defer wg.Done()
		var r UserSite
		site, err1 := Site.GetSiteFromMock()
		if err1 != nil {
			apiError := &apiErrors.ApiError{
				Message: "Fatal URL Site",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		r.Site = site
		c <- r
	}()

	go func() {
		defer wg.Done()
		var r UserSite
		explosure, err2 := Explosure.GetExpFromMock()
		if err2 != nil {
			apiError := &apiErrors.ApiError{
				Message: "Fatal URL Exp",
				Status:  http.StatusBadRequest}
			context.JSON(apiError.Status, apiError)
			return
		}
		r.Explosure = explosure
		c <- r
	}()
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:

			if r.Site != nil {
				result.Site = r.Site
				continue
			}
			if r.Explosure != nil {
				result.Explosure = r.Explosure
				continue
			}
			if r.User != nil {
				result.User = r.User
				continue
			}
		}
	}
	wg.Wait()
	context.JSON(200, &result)
}
