package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPage(t *testing.T) {
	baseURL := "http://localhost:3000"

	var tests = []struct {
		method   string
		url      string
		expected int
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 404},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/6", 200},
		{"GET", "/articles/6/edit", 200},
		{"POST", "/articles/6", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/3/delete", 404},
	}

	for _, test := range tests {
		t.Logf("The current request URL is: %v \n", test.url)

		var (
			resp *http.Response
			err  error
		)

		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseURL+test.url, data)
		default:
			resp, err = http.Get(baseURL + test.url)
		}

		assert.NoError(t, err, "An error occurred when requesting the path:"+test.url)
		assert.Equal(t, test.expected, resp.StatusCode, test.url+" should be return a status code: "+strconv.Itoa(test.expected))
	}
}
