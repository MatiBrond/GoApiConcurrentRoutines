package User

import(
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Utils/ApiErrors"
	"io/ioutil"
	"net/http"
	User2 "github.com/mercadolibre/GoApiConcurrentRoutines/src/api/Domain/User"
)

const urlUsers = "https://api.mercadolibre.com/users/"

func GetUserFromApi(id int64) (*User2.User, *ApiErrors.ApiError){

	var data []byte
	var user User2.User

	urlFinal := fmt.Sprintf("%s%d", urlUsers, id)
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

	if err1 := json.Unmarshal([]byte(data), &user); err1 != nil{
		return nil, &ApiErrors.ApiError{
			Message: "Id is empty user",
			Status: http.StatusBadRequest}
	}
	return &user, nil
}
