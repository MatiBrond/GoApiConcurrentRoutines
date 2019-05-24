package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/services/site"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/utils/apiErrors"
)

const (
	siteID = "id"
)

func GetSiteFromApiCB(context *gin.Context) {

	//siteID := context.Param(siteID)

	site, err1 := site.GetSiteFromApiCBFEDE()
	if err1 != nil {
		apiError := &apiErrors.ApiError{
			Message: err1.Message,
			Status:  http.StatusBadRequest}
		context.JSON(apiError.Status, apiError)
		return
	}
	result := site

	context.JSON(200, &result)

}
