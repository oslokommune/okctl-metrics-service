package meta

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/config"

	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints"
)

// GetRoutes returns meta functionality routes, like the specification
func GetRoutes(specification []byte) endpoints.Routes {
	return endpoints.Routes{
		endpoints.Route{
			Name:        "GetSpecification",
			Method:      http.MethodGet,
			Pattern:     "/specification",
			HandlerFunc: generateGetSpecificationHandler(specification),
		},
	}
}

// GenerateServiceMetaHandler creates a handler for the discovery endpoint, think oath2's .well-known
func GenerateServiceMetaHandler(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, ServiceMeta{
			VersionPrefix: "/v1",
			Specification: fmt.Sprintf("%s/v1/z/specification", getBaseURL(cfg)),
		})
	}
}

// GenerateHealthHandler creates a handler for health checks
func GenerateHealthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
