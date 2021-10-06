package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oslokommune/okctl-metrics-service/pkg/endpoints"
)

func attachRoutes(group *gin.RouterGroup, routes endpoints.Routes) {
	for _, route := range routes {
		attachRoute(group, route)
	}
}

func attachRoute(group *gin.RouterGroup, route endpoints.Route) {
	switch route.Method {
	case http.MethodGet:
		group.GET(route.Pattern, route.HandlerFunc)
	case http.MethodPost:
		group.POST(route.Pattern, route.HandlerFunc)
	case http.MethodPut:
		group.PUT(route.Pattern, route.HandlerFunc)
	case http.MethodPatch:
		group.PATCH(route.Pattern, route.HandlerFunc)
	case http.MethodDelete:
		group.DELETE(route.Pattern, route.HandlerFunc)
	}
}
