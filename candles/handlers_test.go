package candles_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/HackerRank/candles"
)

func TestHandlers(t *testing.T) {
	t.Skip()
	cases := []struct {
		name         string
		given        io.Reader
		responseCode int
	}{
		{
			"Normal fully populated request",
			strings.NewReader(`[1,2,3,4]`),
			http.StatusOK,
		},
		{
			"Empty list returns Bad Request",
			strings.NewReader(`[]`),
			http.StatusBadRequest,
		},
	}

	handler := candles.NewHandler()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "http://localhost:8000", tc.given)
			w := httptest.NewRecorder()
			handler(w, req)

			assert.Equal(t, tc.responseCode, w.Code)
		})
	}
}
