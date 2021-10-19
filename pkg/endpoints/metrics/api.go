package metrics

import (
	"net/http"

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
