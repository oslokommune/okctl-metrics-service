package metrics

const (
	// CategoryCommandExecution represents the context of running commands
	CategoryCommandExecution Category = "commandexecution"

	// ActionScaffoldCluster represents running the command `okctl scaffold cluster`
	ActionScaffoldCluster Action = "scaffoldcluster"
	// ActionApplyCluster represents running the command `okctl apply cluster`
	ActionApplyCluster Action = "applycluster"
	// ActionDeleteCluster represents running the command `okctl delete cluster`
	ActionDeleteCluster Action = "deletecluster"

	// ActionScaffoldApplication represents running the command `okctl scaffold application`
	ActionScaffoldApplication Action = "scaffoldapplication"
	// ActionApplyApplication represents running the command `okctl apply application`
	ActionApplyApplication Action = "applyapplication"

	// ActionForwardPostgres represents running the command `okctl forward postgres`
	ActionForwardPostgres Action = "forwardpostgres"
	// ActionAttachPostgres represents running the command `okctl attach postgres`
	ActionAttachPostgres Action = "attachpostgres"

	// ActionShowCredentials represents running the command `okctl show credentials`
	ActionShowCredentials Action = "showcredentials"
	// ActionUpgrade represents running the command `okctl upgrade`
	ActionUpgrade Action = "upgradecluster"
	// ActionVenv represents running the command `okctl venv`
	ActionVenv Action = "venvcluster"
	// ActionVersion represents running the command `okctl version`
	ActionVersion Action = "version"
)

var commandExecutionActions = []Action{
	ActionScaffoldCluster, ActionApplyCluster, ActionDeleteCluster,
	ActionScaffoldApplication, ActionApplyApplication,
	ActionForwardPostgres, ActionAttachPostgres,
	ActionShowCredentials, ActionUpgrade, ActionVenv, ActionVersion,
}
