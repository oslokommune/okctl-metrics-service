package metrics

// CategoryInstallation represents the context of installing okctl
const CategoryInstallation Category = "installation"

const (
	// ActionInstall represents installing okctl
	ActionInstall Action = "install"
	// ActionBrewUninstall represents uninstalling okctl from brew
	ActionBrewUninstall = "brewuninstall"
)

var installationDefinition = Definition{
	Category: CategoryInstallation,
	Actions: []Action{
		ActionInstall, ActionBrewUninstall,
	},
	Labels: []string{LabelPhaseKey},
}
