package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics"

	"github.com/stretchr/testify/assert"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"
	"github.com/oslokommune/okctl-metrics-service/pkg/router"
)

func TestMetricsStatusCodes(t *testing.T) {
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
					fmt.Sprintf("%s/v1/metrics/events", mockBaseURL),
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
					fmt.Sprintf("%s/v1/metrics/events", mockBaseURL),
					bytes.NewBufferString(`{"category": "commandexecution", "action": "applycluster"}`),
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
					fmt.Sprintf("%s/v1/metrics/events", mockBaseURL),
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
					fmt.Sprintf("%s/v1/metrics/events", mockBaseURL),
					bytes.NewBufferString(`{"category": "commandexecution", "action": "applYcluster"}`),
				)
				req.Header.Add("User-Agent", "okctl")

				return req
			}(),
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name: "Should return 400 upon illegal characters in label",

			withRequest: func() *http.Request {
				req, _ := http.NewRequest(
					http.MethodPost,
					fmt.Sprintf("%s/v1/metrics/events", mockBaseURL),
					bytes.NewBufferString(`{"category": "cluster", "action": "apply", "label": "nese_; DROP ALL TABLES;"}`),
				)
				req.Header.Add("User-Agent", "okctl")

				return req
			}(),
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name: "Should return 400 upon illegal characters in label (less insane)",

			withRequest: func() *http.Request {
				req, _ := http.NewRequest(
					http.MethodPost,
					fmt.Sprintf("%s/v1/metrics/events", mockBaseURL),
					bytes.NewBufferString(`{"category": "cluster", "action": "apply", "label": "test%"}`),
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
			serviceRouter := router.New(generateMockConfig(), generateDemoLogger(), []byte(""))

			recorder := httptest.NewRecorder()

			serviceRouter.ServeHTTP(recorder, tc.withRequest)

			assert.Equal(t, tc.expectStatusCode, recorder.Code)
		})
	}
}

type hit struct {
	Key   string
	Value int
}

func publishEvent(t *testing.T, baseURL string, event metrics.Event) {
	payload, err := json.Marshal(event)
	assert.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/v1/metrics/events", baseURL),
		bytes.NewReader(payload),
	)

	req.Header.Add("User-Agent", "okctl")
	req.Header.Add("Content-Type", "application/json")

	req.RequestURI = ""

	client := http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)

	defer func() {
		_ = res.Body.Close()
	}()

	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func getCounterValue(t *testing.T, baseURL string, key string) int {
	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/z/prometheus", baseURL),
		nil,
	)

	req.RequestURI = ""

	client := http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)

	defer func() {
		_ = res.Body.Close()
	}()

	rawBody, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	re, err := regexp.Compile(fmt.Sprintf("%s (?P<counter>\\d+)", key))
	assert.NoError(t, err)

	result := re.FindStringSubmatch(string(rawBody))
	if len(result) == 0 {
		return 0
	}

	counter, err := strconv.Atoi(result[re.SubexpIndex("counter")])
	assert.NoError(t, err)

	return counter
}

func TestAtoB(t *testing.T) {
	/*
		Due to the Prometheus lib having a global state and no reset/clean functionality, resetting between tests is a
		hassle. I chose to test the difference instead
	*/
	testCases := []struct {
		name       string
		withEvents []metrics.Event
		expectHit  hit
	}{
		{
			name: "Should add and bump metric with one hit",
			withEvents: []metrics.Event{
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionScaffoldCluster,
				},
			},
			expectHit: hit{
				Key:   "okctl_commandexecution_scaffoldcluster",
				Value: 1,
			},
		},
		{
			name: "Should add and bump metric with multiple hits",
			withEvents: []metrics.Event{
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionScaffoldCluster,
				},
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionScaffoldCluster,
				},
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionScaffoldCluster,
				},
			},
			expectHit: hit{
				Key:   "okctl_commandexecution_scaffoldcluster",
				Value: 3,
			},
		},
		{
			name: "Should handle labels",
			withEvents: []metrics.Event{
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionShowCredentials,
					Labels: map[string]string{
						"phase": "start",
					},
				},
			},
			expectHit: hit{
				Key:   `okctl_commandexecution_showcredentials{phase="start"}`,
				Value: 1,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(router.New(generateMockConfig(), generateDemoLogger(), []byte("")))
			defer server.Close()

			originalValue := getCounterValue(t, server.URL, tc.expectHit.Key)

			for _, event := range tc.withEvents {
				publishEvent(t, server.URL, event)
			}

			newValue := getCounterValue(t, server.URL, tc.expectHit.Key)

			diff := math.Abs(float64(newValue - originalValue))

			assert.Equal(t, tc.expectHit.Value, int(diff))
		})
	}
}

const mockBaseURL = "http://localhost:3000"

func generateMockConfig() config.Config {
	cfg, _ := config.Generate()

	cfg.BaseURL = "http://localhost"
	cfg.Port = 3000

	return cfg
}

func generateDemoLogger() *logrus.Logger {
	logger := logrus.New()

	logger.Out = io.Discard

	return logger
}
