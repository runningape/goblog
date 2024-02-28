package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	baseURL := "http://localhost:3000"

	var (
		resp *http.Response
		err  error
	)
	resp, err = http.Get(baseURL + "/")

	assert.NoError(t, err, "An error has occurred, err is not nil.")
	assert.Equal(t, 200, resp.StatusCode, "It should return a status code of 200.")
}
