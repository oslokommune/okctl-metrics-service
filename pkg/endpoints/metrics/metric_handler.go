package metrics

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/core"
)

var counters = make(map[string]prometheus.Counter)

func generateMetricHandler(cfg config.Config, logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")

		err := validateAgent(cfg.LegalAgents, userAgent)
		if err != nil {
			err = fmt.Errorf("validating agent %s: %w", userAgent, err)

			c.JSON(http.StatusForbidden, core.ErrorResponse{Error: err.Error()})

			logger.Warnf("invalid user agent '%s'. Legal agents are %s", userAgent, cfg.LegalAgents)

			return
		}

		event := NewEvent()

		err = c.Bind(&event)
		if err != nil {
			c.Status(http.StatusBadRequest)

			logger.Debug("binding request body: ", err.Error())

			return
		}

		err = event.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, core.ErrorResponse{Error: err.Error()})

			logger.Debug("invalid event: ", err.Error())

			return
		}

		err = registerMetric(userAgent, event)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			logger.Error("registering metric: ", err.Error())

			return
		}

		c.Status(http.StatusCreated)
	}
}

func registerMetric(prefix string, event Event) error {
	key := event.Hash()

	if _, ok := counters[key]; !ok {
		counters[key] = prometheus.NewCounter(prometheus.CounterOpts{
			Namespace:   prefix,
			Subsystem:   event.Category.String(),
			Name:        event.Action.String(),
			ConstLabels: event.Labels,
		})

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
