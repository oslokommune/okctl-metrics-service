package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/config"
	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints"
	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/meta"
	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics"
	"github.com/sirupsen/logrus"
)

// New configures a new router for handling requests to the service
func New(cfg config.Config, logger *logrus.Logger, specification []byte) (*gin.Engine, endpoints.TeardownFn) {
	router := gin.New()
	teardowns := make([]endpoints.TeardownFn, 0)

	configureLogging(router, logger)

	configureMetaRoutes(router, cfg)

	teardowns = append(teardowns, configureV1Routes(router, cfg, logger, specification)...)

	return router, func() {
		for _, teardownFn := range teardowns {
			teardownFn()
		}
	}
}

func configureLogging(router *gin.Engine, logger *logrus.Logger) {
	router.Use(gin.Recovery())

	skipPaths := []string{
		"/",
		"/z/health",
		"/z/ready",
		"/z/prometheus",
	}

	router.Use(generateJSONLoggerMiddleware(logger, skipPaths))
}

func configureMetaRoutes(router *gin.Engine, cfg config.Config) {
	router.GET("/", meta.GenerateServiceMetaHandler(cfg))
	router.GET("/z/health", meta.GenerateHealthHandler())
	router.GET("/z/ready", meta.GenerateHealthHandler())
	router.GET("/z/prometheus", meta.GeneratePrometheusHandler())
}

func configureV1Routes(router *gin.Engine, cfg config.Config, logger *logrus.Logger, specification []byte) []endpoints.TeardownFn {
	teardowns := make([]endpoints.TeardownFn, 0)
	v1Group := router.Group("/v1")

	v1MetaGroup := v1Group.Group("/z")
	attachRoutes(v1MetaGroup, meta.GetRoutes(specification))

	v1MetricsGroup := v1Group.Group("/metrics")
	v1MetricsRoutes, v1MetricsTeardowns := metrics.GetRoutes(cfg, logger)

	attachRoutes(v1MetricsGroup, v1MetricsRoutes)
	teardowns = append(teardowns, v1MetricsTeardowns...)

	return teardowns
}
