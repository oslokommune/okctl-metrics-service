package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/config"
	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/meta"
	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics"
)

// New configures a new router for handling requests to the service
func New(cfg config.Config, specification []byte) *gin.Engine {
	router := gin.New()

	configureLogging(router)

	configureMetaRoutes(router, cfg)

	configureV1Routes(router, cfg, specification)

	return router
}

func configureLogging(router *gin.Engine) {
	router.Use(gin.Recovery())

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{
			"/",
			"/z/health",
			"/z/ready",
			"/z/prometheus",
		},
	}))
}

func configureMetaRoutes(router *gin.Engine, cfg config.Config) {
	router.GET("/", meta.GenerateServiceMetaHandler(cfg))
	router.GET("/z/health", meta.GenerateHealthHandler())
	router.GET("/z/ready", meta.GenerateHealthHandler())
	router.GET("/z/prometheus", meta.GeneratePrometheusHandler())
}

func configureV1Routes(router *gin.Engine, cfg config.Config, specification []byte) {
	v1Group := router.Group("/v1")

	v1MetaGroup := v1Group.Group("/z")
	attachRoutes(v1MetaGroup, meta.GetRoutes(specification))

	v1MetricsGroup := v1Group.Group("/metrics")
	attachRoutes(v1MetricsGroup, metrics.GetRoutes(cfg))
}
