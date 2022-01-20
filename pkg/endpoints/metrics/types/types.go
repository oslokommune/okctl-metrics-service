package types

type (
	// Category represents the category of which an event belongs to
	Category string
	// Action represents the action in the context of a Category has been executed
	Action string
)

// Definition represents the definition of a metric
type Definition struct {
	Category Category
	Actions  []Action
	Labels   []string
}
