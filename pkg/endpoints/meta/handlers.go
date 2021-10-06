package meta

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func generateGetSpecificationHandler(specification []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Data(http.StatusOK, "application/x-yaml", specification)
	}
}
