package circuitBraker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mercadolibre/GoApiConcurrentRoutines/src/api/utils/apiErrors"
)

type State int

const (
	StateClosed State = iota
	StateHalfOpen
	StateOpen
)

type TimeControl struct {
	timeOut     time.Duration
	timeCurrent time.Time
	countFail   int64
}

type CircuitBreaker struct {
	state   State
	control TimeControl
}

func NewCB() *CircuitBreaker {
	return &CircuitBreaker{
		StateClosed,
		TimeControl{
			0,
			time.Now(),
			0,
		},
	}
}

func (cb *CircuitBreaker) GetCircuitBreaker(url string) ([]byte, *apiErrors.ApiError) {

	switch cb.state {

	case StateOpen:

		cb.control.countFail += 1

		if (time.Now().Second()-cb.control.timeCurrent.Second()) > 15 || (time.Now().Second()-cb.control.timeCurrent.Second()) < -15 {
			cb.state = StateHalfOpen
		}

		return nil, &apiErrors.ApiError{
			"La API se encuentra caída.",
			http.StatusBadRequest,
		}

	case StateHalfOpen:

		res, err := http.Get(url)
		if err != nil || res.StatusCode == 500 {

			cb.state = StateOpen
			cb.control.countFail += 1
			cb.control.timeCurrent = time.Now()

			return nil, &apiErrors.ApiError{
				"La Api continúa caída",
				500,
			}

		} else {

			cb.state = StateClosed
			cb.control.countFail = 0

			data, error := ioutil.ReadAll(res.Body)

			if error != nil {
				return nil, &apiErrors.ApiError{
					error.Error(),
					http.StatusBadRequest,
				}
			}
			return data, nil
		}
	}

	res, err := http.Get(url)

	if (err != nil) || (res.StatusCode == 500) {

		cb.control.countFail += 1

		if cb.control.countFail >= 3 {
			cb.state = StateOpen
			cb.control.timeCurrent = time.Now()
			fmt.Println("3 fallas seguidas")
		}
		fmt.Println("falla")

		return nil, &apiErrors.ApiError{
			"La Api esta caída",
			500,
		}

	} else {

		cb.state = StateClosed
		cb.control.countFail = 0
		data, error := ioutil.ReadAll(res.Body)

		if error != nil {
			return nil, &apiErrors.ApiError{
				error.Error(),
				http.StatusBadRequest,
			}
		}
		return data, nil
	}

}
