package installation

import "github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics/types"

// Category represents the context of installing okctl
const Category types.Category = "installation"

const (
	// ActionInstall represents installing okctl
	ActionInstall types.Action = "install"
)

const (
	// LabelPhaseKey is the key for phase label
	LabelPhaseKey = "phase"
	// LabelPhaseStart represents the start of a command
	LabelPhaseStart = "start"
	// LabelPhaseEnd represents the end of the command
	LabelPhaseEnd = "end"
)

var Definition = types.Definition{
	Category: Category,
	Actions: []types.Action{
		ActionInstall,
	},
	Labels: []string{LabelPhaseKey},
}
