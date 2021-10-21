package metrics

// Event defines information necessary to process a metric event
type Event struct {

	// A label used to categorize events
	Category Category `json:"category"`

	// A label used to identify the event type
	Action Action `json:"action"`

	// Labels that annotate the event
	Labels map[string]string `json:"labels"`
}

type (
	Category string
	Action   string
)

func (c Category) String() string {
	return string(c)
}

func (a Action) String() string {
	return string(a)
}
