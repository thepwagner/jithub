package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thepwagner/jithub/api"
)

func TestHandler_ServeHTTP_404(t *testing.T) {
	cases := []string{
		"/",
		"/robots.txt",
	}
	for _, tc := range cases {
		t.Run(tc, func(t *testing.T) {
			req, err := http.NewRequest("HEAD", tc, nil)
			require.NoError(t, err)
			res := &httptest.ResponseRecorder{}

			h := api.Handler{}
			h.ServeHTTP(res, req)

			assert.Equal(t, http.StatusNotFound, res.Code)
		})
	}
}

func TestHandler_ServeHTTP(t *testing.T) {
	req, err := http.NewRequest("HEAD", "/com/github/thepwagner/test-java/1.0/test-java-1.0.pom", nil)
	require.NoError(t, err)
	res := &httptest.ResponseRecorder{}

	h := api.Handler{}
	h.ServeHTTP(res, req)
}
