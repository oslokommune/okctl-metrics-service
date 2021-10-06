package metrics

// Event defines information necessary to process a metric event
type Event struct {

	// A UUID used to combine events
	ID string `json:"ID,omitempty"`

	// A label used to categorize events
	Category string `json:"Category,omitempty"`

	// A label used to identify the event type
	Action string `json:"Action,omitempty"`
}
