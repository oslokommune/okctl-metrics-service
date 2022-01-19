package metrics

import "github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics/types"

const defaultMetricNamespace = "okctl"

// Event defines information necessary to process a metric event
type Event struct {

	// A label used to categorize events
	Category types.Category `json:"category"`

	// A label used to identify the event type
	Action types.Action `json:"action"`

	// Labels that annotate the event
	Labels map[string]string `json:"labels"`
}
