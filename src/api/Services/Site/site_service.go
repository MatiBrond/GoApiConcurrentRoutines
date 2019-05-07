package Site

import (
	"encoding/json"
	Site2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Domain/Site"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Utils/ApiErrors"
	"io/ioutil"
	"net/http"
)

const urlSite = "https://api.mercadolibre.com/sites/"

func GetSiteFromApi( id string) (*Site2.Site, *ApiErrors.ApiError){

	var data []byte
	var site Site2.Site
	urlFinal := urlSite + id
	response, err := http.Get(urlFinal)

	if err != nil {
		return nil, &ApiErrors.ApiError{
			Message: "Fatal URL",
			Status: http.StatusBadRequest}
	}
	data, error := ioutil.ReadAll(response.Body)

	if error != nil{
		return nil, &ApiErrors.ApiError{
			Message: "Fallo aca",
			Status: http.StatusInternalServerError}
	}

	if err1 := json.Unmarshal([]byte(data), &site); err1 != nil{
		return nil, &ApiErrors.ApiError{
			Message: "Id is empty site",
			Status: http.StatusBadRequest}
	}
	return &site, nil
}