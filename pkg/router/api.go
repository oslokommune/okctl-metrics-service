package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/config"
	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/meta"
	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics"
)

// New configures a new router for handling requests to the service
func New(cfg config.Config, specification []byte) *gin.Engine {
	router := gin.Default()

	router.GET("/", meta.GenerateServiceMetaHandler(cfg))
	router.GET("/z/health", meta.GenerateHealthHandler())
	router.GET("/z/ready", meta.GenerateHealthHandler())
	router.GET("/z/prometheus", meta.GeneratePrometheusHandler())

	configureV1Routes(router, specification)

	return router
}

func configureV1Routes(router *gin.Engine, specification []byte) {
	v1Group := router.Group("/v1")

	v1MetaGroup := v1Group.Group("/z")
	attachRoutes(v1MetaGroup, meta.GetRoutes(specification))

	v1MetricsGroup := v1Group.Group("/metrics")
	attachRoutes(v1MetricsGroup, metrics.GetRoutes())
}
