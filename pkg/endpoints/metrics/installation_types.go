package metrics

// CategoryInstallation represents the context of installing okctl
const CategoryInstallation Category = "installation"

const (
	// ActionInstall represents installing okctl
	ActionInstall Action = "install"
)

var installationDefinition = Definition{
	Category: CategoryInstallation,
	Actions: []Action{
		ActionInstall,
	},
	Labels: []string{LabelPhaseKey},
}
