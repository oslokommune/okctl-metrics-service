package metrics

// Event defines information necessary to process a metric event
type Event struct {

	// A label used to categorize events
	Category Category `json:"category"`

	// A label used to identify the event type
	Action Action `json:"action"`
}

type (
	Category string
	Action   string
)

// Categories
const (
	// CategoryCluster represents metrics associated with cluster manipulation
	CategoryCluster Category = "cluster"
	// CategoryApplication represents metrics associated with application manipulation
	CategoryApplication Category = "application"
)

// Actions
const (
	// ActionScaffold represents scaffolding a resource
	ActionScaffold Action = "scaffold"
	// ActionApply represents applying a resource
	ActionApply Action = "apply"
	// ActionDelete represents deleting a resource
	ActionDelete Action = "delete"
)
