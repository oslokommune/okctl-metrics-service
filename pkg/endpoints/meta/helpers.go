package meta

import (
	"fmt"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"
)

func getBaseURL(cfg config.Config) string {
	return fmt.Sprintf("%s:%d", cfg.BaseURL, cfg.Port)
}
