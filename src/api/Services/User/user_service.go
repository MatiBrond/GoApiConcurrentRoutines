package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	User2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/domain/user"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/utils/apiErrors"
)

const urlUsers = "https://api.mercadolibre.com/users/"
const localURL = "http://localhost:8089/users/"

func GetUserFromApi(id int64) (*User2.User, *apiErrors.ApiError) {

	var data []byte
	var user User2.User

	urlFinal := fmt.Sprintf("%s%d", urlUsers, id)
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

	if err1 := json.Unmarshal([]byte(data), &user); err1 != nil {
		return nil, &apiErrors.ApiError{
			Message: "Id is empty user",
			Status:  http.StatusBadRequest}
	}
	return &user, nil
}

func GetUserFromMock() (*User2.User, *apiErrors.ApiError) {

	var data []byte
	var user User2.User

	response, err := http.Get(localURL)

	if err != nil {
		return nil, &apiErrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest}

	}
	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, &apiErrors.ApiError{
			Message: "Fallo aca",
			Status:  http.StatusInternalServerError}
	}

	if err1 := json.Unmarshal([]byte(data), &user); err1 != nil {
		return nil, &apiErrors.ApiError{
			Message: "Id is empty user",
			Status:  http.StatusBadRequest}
	}
	return &user, nil
}
