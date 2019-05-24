package site

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	circuit_breaker "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/circuitBraker"
	Site2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/domain/site"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/utils/apiErrors"
	"github.com/sony/gobreaker"
)

const urlSite = "https://api.mercadolibre.com/sites/"
const localURL = "http://localhost:8089/sites/"
const mockURL = "http://localhost:8085/site/"

var cb *gobreaker.CircuitBreaker

var cbFede circuit_breaker.CircuitBreaker

func getCBreaker() *circuit_breaker.CircuitBreaker {
	if &cbFede != nil {
		return &cbFede
	}
	return circuit_breaker.NewCB()
}

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	st.Timeout = 200 * time.Millisecond
	st.Interval = 200 * time.Millisecond

	cb = gobreaker.NewCircuitBreaker(st)
}

func GetSiteFromApiCBFEDE() (*Site2.Site, *apiErrors.ApiError) {

	var site Site2.Site
	//urlFinal := urlSite + id
	cbFede := getCBreaker()
	//Le paso la responsabilidad de la request al gobreaker
	body, err := cbFede.GetCircuitBreaker(mockURL)

	if err != nil {
		return nil, &apiErrors.ApiError{
			Message: err.Message,
			Status:  err.Status}
	}

	if err1 := json.Unmarshal([]byte(body), &site); err1 != nil {
		return nil, &apiErrors.ApiError{
			Message: err1.Error(),
			Status:  http.StatusBadRequest}
	}
	return &site, nil

}

// func GetSiteFromApiCB() (*Site2.Site, *ApiErrors.ApiError) {

// 	var site Site2.Site
// 	//urlFinal := urlSite + id

// 	//Le paso la responsabilidad de la request al gobreaker
// 	body, err := cb.Execute(func() (interface{}, error) {

// 		resp, err := http.Get(mockURL)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer resp.Body.Close()
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return body, nil

// 	})
// 	if err != nil {
// 		return nil, &ApiErrors.ApiError{
// 			Message: err.Error(),
// 			Status:  http.StatusBadRequest}
// 	}

// 	body2 := body.([]byte) //Cambio el body (interface{}) -> body2 ([]byte)

// 	if err1 := json.Unmarshal([]byte(body2), &site); err1 != nil {
// 		return nil, &ApiErrors.ApiError{
// 			Message: "Id is empty site",
// 			Status:  http.StatusBadRequest}
// 	}

// 	return &site, nil

// }

func GetSiteFromApi(id string) (*Site2.Site, *apiErrors.ApiError) {

	var data []byte
	var site Site2.Site
	urlFinal := urlSite + id

	response, err := http.Get(urlFinal)

	if err != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fatal URLLL",
			Status:  http.StatusBadRequest}
	}
	data, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fallo aca",
			Status:  http.StatusInternalServerError}
	}

	if err1 := json.Unmarshal([]byte(data), &site); err1 != nil {
		return nil, &apiErrors.ApiError{
			Message: "Id is empty site",
			Status:  http.StatusBadRequest}
	}
	return &site, nil
}

func GetSiteFromMock() (*Site2.Site, *apiErrors.ApiError) {

	var data []byte
	var site Site2.Site
	response, err := http.Get(localURL)

	if err != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fatal URURRRRLL",
			Status:  http.StatusBadRequest}
	}
	data, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fallo aca",
			Status:  http.StatusInternalServerError}
	}

	if err1 := json.Unmarshal([]byte(data), &site); err1 != nil {
		return nil, &apiErrors.ApiError{
			Message: "Id is empty site",
			Status:  http.StatusBadRequest}
	}
	return &site, nil
}
