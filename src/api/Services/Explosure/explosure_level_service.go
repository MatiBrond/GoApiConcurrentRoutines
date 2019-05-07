package Explosure

import (
	"encoding/json"
	Explosure2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Domain/Explosure"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Utils/ApiErrors"
	"io/ioutil"
	"net/http"
)

const urlSite = "https://api.mercadolibre.com/sites/"

func GetUserFromApi(id string) (*Explosure2.Explosure_level, *ApiErrors.ApiError){

	var data []byte
	var explosure_level Explosure2.Explosure_level
	urlFinal := urlSite + id + "/listing_exposures"

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

	if err1 := json.Unmarshal([]byte(data), &explosure_level); err1 != nil{
		return nil, &ApiErrors.ApiError{
			Message: "Id is empty explosure",
			Status: http.StatusBadRequest}
	}
	return &explosure_level, nil
}
