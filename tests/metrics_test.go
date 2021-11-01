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

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics"

	"github.com/stretchr/testify/assert"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"
	"github.com/oslokommune/okctl-metrics-service/pkg/router"
)

func TestMetricsStatusCodes(t *testing.T) {
	testCases := []struct {
		name string

		withEvent        metrics.Event
		withUserAgent    string
		expectStatusCode int
	}{
		{
			name: "Should return 403 upon erroneous user agent",
			withEvent: metrics.Event{
				Category: metrics.CategoryCommandExecution,
				Action:   metrics.ActionVenv,
				Labels: map[string]string{
					metrics.LabelPhaseKey: metrics.LabelPhaseStart,
				},
			},
			withUserAgent:    "chromium",
			expectStatusCode: http.StatusForbidden,
		},
		{
			name: "Should return 201 upon expected request",
			withEvent: metrics.Event{
				Category: metrics.CategoryCommandExecution,
				Action:   metrics.ActionApplyCluster,
				Labels: map[string]string{
					metrics.LabelPhaseKey: metrics.LabelPhaseStart,
				},
			},
			withUserAgent:    mockLegalUserAgent,
			expectStatusCode: http.StatusCreated,
		},
		{
			name: "Should return 400 upon unexpected category",

			withEvent: metrics.Event{
				Category: "automation",
				Action:   metrics.ActionVenv,
				Labels:   map[string]string{metrics.LabelPhaseKey: metrics.LabelPhaseStart},
			},
			withUserAgent:    mockLegalUserAgent,
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name: "Should return 400 upon unexpected action",
			withEvent: metrics.Event{
				Category: metrics.CategoryCommandExecution,
				Action:   "applYcluster",
				Labels:   map[string]string{metrics.LabelPhaseKey: metrics.LabelPhaseStart},
			},
			withUserAgent:    mockLegalUserAgent,
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name: "Should return 400 upon illegal characters in label",
			withEvent: metrics.Event{
				Category: metrics.CategoryCommandExecution,
				Action:   metrics.ActionApplyCluster,
				Labels:   map[string]string{metrics.LabelPhaseKey: "nese_; DROP ALL TABLES;"},
			},
			withUserAgent:    mockLegalUserAgent,
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name: "Should return 400 upon illegal characters in label (less insane)",
			withEvent: metrics.Event{
				Category: metrics.CategoryCommandExecution,
				Action:   metrics.ActionApplyCluster,
				Labels:   map[string]string{metrics.LabelPhaseKey: "test%"},
			},
			withUserAgent:    mockLegalUserAgent,
			expectStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			serviceRouter, teardownFn := router.New(generateMockConfig(), generateDemoLogger(), []byte(""))
			defer teardownFn()

			recorder := httptest.NewRecorder()

			payload, err := json.Marshal(tc.withEvent)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				fmt.Sprintf("%s/v1/metrics/events", mockBaseURL),
				bytes.NewReader(payload),
			)
			assert.NoError(t, err)

			req.Header.Add("User-Agent", tc.withUserAgent)
			req.Header.Add("Content-Type", "application/json")

			serviceRouter.ServeHTTP(recorder, req)

			assert.Equal(t, tc.expectStatusCode, recorder.Code)
		})
	}
}

type hit struct {
	Key   string
	Value int
}

func getCounterValue(t *testing.T, server *gin.Engine, key string) int {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/z/prometheus", mockBaseURL), nil)
	assert.NoError(t, err)

	server.ServeHTTP(recorder, req)

	rawBody, err := io.ReadAll(recorder.Body)
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
					Labels:   map[string]string{metrics.LabelPhaseKey: metrics.LabelPhaseStart},
				},
			},
			expectHit: hit{
				Key:   `okctl_commandexecution_scaffoldcluster{phase="start"}`,
				Value: 1,
			},
		},
		{
			name: "Should add and bump metric with multiple hits",
			withEvents: []metrics.Event{
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionScaffoldCluster,
					Labels:   map[string]string{metrics.LabelPhaseKey: metrics.LabelPhaseStart},
				},
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionScaffoldCluster,
					Labels:   map[string]string{metrics.LabelPhaseKey: metrics.LabelPhaseStart},
				},
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionScaffoldCluster,
					Labels:   map[string]string{metrics.LabelPhaseKey: metrics.LabelPhaseStart},
				},
			},
			expectHit: hit{
				Key:   `okctl_commandexecution_scaffoldcluster{phase="start"}`,
				Value: 3,
			},
		},
		{
			name: "Should handle labels",
			withEvents: []metrics.Event{
				{
					Category: metrics.CategoryCommandExecution,
					Action:   metrics.ActionShowCredentials,
					Labels:   map[string]string{metrics.LabelPhaseKey: metrics.LabelPhaseStart},
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
			serviceRouter, teardownFn := router.New(generateMockConfig(), generateDemoLogger(), []byte(""))
			defer teardownFn()

			originalValue := getCounterValue(t, serviceRouter, tc.expectHit.Key)

			for _, event := range tc.withEvents {
				publishEvent(t, serviceRouter, mockLegalUserAgent, event)
			}

			newValue := getCounterValue(t, serviceRouter, tc.expectHit.Key)

			diff := math.Abs(float64(newValue - originalValue))

			assert.Equal(t, tc.expectHit.Value, int(diff))
		})
	}
}

const (
	mockBaseURL        = "http://localhost:3000"
	mockLegalUserAgent = "okctl"
)

func publishEvent(t *testing.T, serviceRouter *gin.Engine, userAgent string, event metrics.Event) {
	payload, err := json.Marshal(event)
	assert.NoError(t, err)

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/v1/metrics/events", mockBaseURL),
		bytes.NewReader(payload),
	)
	assert.NoError(t, err)

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	serviceRouter.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

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
