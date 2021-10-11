package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"
	"github.com/oslokommune/okctl-metrics-service/pkg/router"
)

func TestMetrics(t *testing.T) {
	testCases := []struct {
		name string

		withRequest      *http.Request
		expectStatusCode int
	}{
		{
			name: "Should return 403 upon erroneous user agent",

			withRequest: func() *http.Request {
				req, _ := http.NewRequest(
					http.MethodPost,
					fmt.Sprintf("%s/metrics/events", mockURL),
					bytes.NewBufferString("{}"),
				)
				req.Header.Add("User-Agent", "chromium")

				return req
			}(),
			expectStatusCode: http.StatusForbidden,
		},
		{
			name: "Should return 201 upon expected request",

			withRequest: func() *http.Request {
				req, _ := http.NewRequest(
					http.MethodPost,
					fmt.Sprintf("%s/metrics/events", mockURL),
					bytes.NewBufferString(`{"category": "cluster", "action": "apply"}`),
				)
				req.Header.Add("User-Agent", "okctl")
				req.Header.Add("Content-Type", "application/json")

				return req
			}(),
			expectStatusCode: http.StatusCreated,
		},
		{
			name: "Should return 400 upon unexpected category",

			withRequest: func() *http.Request {
				req, _ := http.NewRequest(
					http.MethodPost,
					fmt.Sprintf("%s/metrics/events", mockURL),
					bytes.NewBufferString(`{"category": "automation", "action": "apply"}`),
				)
				req.Header.Add("User-Agent", "okctl")

				return req
			}(),
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name: "Should return 400 upon unexpected action",

			withRequest: func() *http.Request {
				req, _ := http.NewRequest(
					http.MethodPost,
					fmt.Sprintf("%s/metrics/events", mockURL),
					bytes.NewBufferString(`{"category": "cluster", "action": "applY"}`),
				)
				req.Header.Add("User-Agent", "okctl")

				return req
			}(),
			expectStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			serviceRouter := router.New(generateMockConfig(), []byte(""))

			recorder := httptest.NewRecorder()

			serviceRouter.ServeHTTP(recorder, tc.withRequest)

			assert.Equal(t, tc.expectStatusCode, recorder.Code)
		})
	}
}

const mockURL = "http://localhost:3000/v1"

func generateMockConfig() config.Config {
	cfg, _ := config.Generate()

	cfg.BaseURL = "http://localhost"
	cfg.Port = 3000

	return cfg
}
