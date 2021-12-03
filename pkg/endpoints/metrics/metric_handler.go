package metrics

import (
	"fmt"
	"net/http"

	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints"

	"github.com/sirupsen/logrus"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/core"
)

func generateMetricHandler(cfg config.Config, logger *logrus.Logger) (gin.HandlerFunc, endpoints.TeardownFn) {
	counters := NewMetricRegistry()

	counters.Add(commandExecutionDefinition)
	counters.Add(installationDefinition)
	counters.Add(brewOkctlInstallationDefinition)

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

			logger.Debug("validating event: ", err.Error())

			return
		}

		err = counters.Increment(userAgent, event)
		if err != nil {
			c.JSON(http.StatusBadRequest, core.ErrorResponse{Error: err.Error()})

			logger.Debug("incrementing metric: ", err.Error())

			return
		}

		c.Status(http.StatusCreated)
	}, counters.Reset
}
