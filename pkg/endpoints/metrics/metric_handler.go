package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/core"
)

const requiredUserAgent = "okctl"

var counters = make(map[string]prometheus.Counter)

func generateMetricHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		if userAgent != requiredUserAgent {
			c.Status(http.StatusForbidden)

			return
		}

		var event Event

		err := c.Bind(&event)
		if err != nil {
			c.Status(http.StatusBadRequest)

			return
		}

		err = event.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, core.ErrorResponse{Error: err.Error()})

			return
		}

		err = registerMetric(event)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		c.Status(http.StatusCreated)
	}
}

func registerMetric(event Event) error {
	key := fmt.Sprintf("okctl_%s_%s", event.Category, event.Action)

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
