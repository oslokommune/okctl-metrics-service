package metrics

import (
	"net/http"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints"
)

// GetRoutes returns endpoints related to metric handling
func GetRoutes(cfg config.Config) endpoints.Routes {
	return endpoints.Routes{
		endpoints.Route{
			Name:        "SubmitEvent",
			Method:      http.MethodPost,
			Pattern:     "/events",
			HandlerFunc: generateMetricHandler(cfg),
		},
	}
}

// Validate ensures an Event contains the required and valid data
func (receiver Event) Validate() error {
	return validation.ValidateStruct(&receiver,
		validation.Field(&receiver.Category, validation.Required),
		validation.Field(&receiver.Action, validation.Required),
	)
}
