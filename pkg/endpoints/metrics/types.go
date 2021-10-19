package metrics

// Event defines information necessary to process a metric event
type Event struct {

	// A label used to categorize events
	Category Category `json:"category"`

	// A label used to identify the event type
	Action Action `json:"action"`

	// A label used to determine variations of an event
	Label string `json:"label"`
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
