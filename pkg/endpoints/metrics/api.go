package metrics

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"

	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints"
)

// GetRoutes returns endpoints related to metric handling
func GetRoutes(cfg config.Config, logger *logrus.Logger) endpoints.Routes {
	return endpoints.Routes{
		endpoints.Route{
			Name:        "SubmitEvent",
			Method:      http.MethodPost,
			Pattern:     "/events",
			HandlerFunc: generateMetricHandler(cfg, logger),
		},
	}
}

// NewEvent initializes a metric event
func NewEvent() Event {
	return Event{Labels: make(map[string]string)}
}

// Hash generates a consequent string representing an event
func (receiver Event) Hash() string {
	hasher := sha256.New()

	parts := []string{
		receiver.Category.String(),
		receiver.Action.String(),
		mapAsString(receiver.Labels),
	}

	_, _ = hasher.Write([]byte(strings.Join(parts, "")))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
