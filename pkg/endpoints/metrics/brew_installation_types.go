package metrics

// CategoryBrewOkctlInstallation represents the context of installing okctl with brew
const CategoryBrewOkctlInstallation Category = "brewokctlinstallation"

const (
	// ActionUninstall represents uninstalling okctl with brew
	ActionUninstall = "uninstall"
)

var brewOkctlInstallationDefinition = Definition{
	Category: CategoryBrewOkctlInstallation,
	Actions: []Action{
		ActionUninstall,
	},
}
