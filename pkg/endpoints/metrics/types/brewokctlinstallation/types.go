package brewokctlinstallation

import "github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics/types"

// Category represents the context of installing okctl with brew
const Category types.Category = "brewokctlinstallation"

const (
	// ActionUninstall represents uninstalling okctl with brew
	ActionUninstall = "uninstall"
)

var Definition = types.Definition{
	Category: Category,
	Actions: []types.Action{
		ActionUninstall,
	},
}
