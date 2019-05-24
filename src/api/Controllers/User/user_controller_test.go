package user

import (
	"net/http"
	"testing"
)

func BenchmarkGetMock(b *testing.B) {

	for i := 0; i < b.N; i++ {

		http.Get("http://localhost:8080/usersmock/")

	}
}
