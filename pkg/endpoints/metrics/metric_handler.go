package metrics

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/core"
)

var counters = make(map[string]prometheus.Counter)

func generateMetricHandler(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")

		err := validateAgent(cfg.LegalAgents, userAgent)
		if err != nil {
			err = fmt.Errorf("validating agent %s: %w", userAgent, err)

			c.JSON(http.StatusForbidden, core.ErrorResponse{Error: err.Error()})

			return
		}

		var event Event

		err = c.Bind(&event)
		if err != nil {
			c.Status(http.StatusBadRequest)

			return
		}

		err = event.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, core.ErrorResponse{Error: err.Error()})

			return
		}

		err = registerMetric(userAgent, event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, core.ErrorResponse{Error: err.Error()})

			return
		}

		c.Status(http.StatusCreated)
	}
}

func registerMetric(prefix string, event Event) error {
	key := eventAsKey(prefix, event)

	if _, ok := counters[key]; !ok {
		counters[key] = prometheus.NewCounter(prometheus.CounterOpts{Name: key})

		err := prometheus.Register(counters[key])
		if err != nil {
			return fmt.Errorf("registering Prometheus counter: %w", err)
		}
	}

	counters[key].Inc()

	return nil
}

func validateAgent(legalAgents []string, agent string) error {
	for _, legalAgent := range legalAgents {
		if agent == legalAgent {
			return nil
		}
	}

	return fmt.Errorf("invalid agent")
}

func eventAsKey(prefix string, event Event) string {
	nameParts := []string{string(event.Action)}

	if event.Label != "" {
		nameParts = append(nameParts, strings.ToLower(event.Label))
	}

	return prometheus.BuildFQName(prefix, string(event.Category), strings.Join(nameParts, "_"))
}
