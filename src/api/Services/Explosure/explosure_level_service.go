package explosure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	Explosure2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/domain/explosure"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/utils/apiErrors"
)

const urlSite = "https://api.mercadolibre.com/sites/"
const localUrl = "http://localhost:8089/exposures/"

func GetExpFromApi(id string) (*Explosure2.Explosure_level, *apiErrors.ApiError) {

	var data []byte
	var explosure_level Explosure2.Explosure_level
	urlFinal := urlSite + id + "/listing_exposures"

	response, err := http.Get(urlFinal)
	if err != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fatal URL",
			Status:  http.StatusBadRequest}
	}
	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fallo aca",
			Status:  http.StatusInternalServerError}
	}

	if err1 := json.Unmarshal([]byte(data), &explosure_level); err1 != nil {
		return nil, &apiErrors.ApiError{
			Message: "Id is empty explosure",
			Status:  http.StatusBadRequest}
	}
	return &explosure_level, nil
}

func GetExpFromMock() (*Explosure2.Explosure_level, *apiErrors.ApiError) {

	var data []byte
	var explosure_level Explosure2.Explosure_level

	response, err := http.Get(localUrl)
	if err != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fatal URL",
			Status:  http.StatusBadRequest}
	}
	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fallo aca",
			Status:  http.StatusInternalServerError}
	}

	if err1 := json.Unmarshal([]byte(data), &explosure_level); err1 != nil {
		return nil, &apiErrors.ApiError{
			Message: "Id is empty explosure",
			Status:  http.StatusBadRequest}
	}
	return &explosure_level, nil
}
