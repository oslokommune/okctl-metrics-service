package commandexecution

import "github.com/oslokommune/okctl-metrics-service/pkg/endpoints/metrics/types"

// Category represents the context of running commands
const Category types.Category = "commandexecution"

const (
	// ActionScaffoldCluster represents running the command `okctl scaffold cluster`
	ActionScaffoldCluster types.Action = "scaffoldcluster"
	// ActionApplyCluster represents running the command `okctl apply cluster`
	ActionApplyCluster types.Action = "applycluster"
	// ActionDeleteCluster represents running the command `okctl delete cluster`
	ActionDeleteCluster types.Action = "deletecluster"

	// ActionScaffoldApplication represents running the command `okctl scaffold application`
	ActionScaffoldApplication types.Action = "scaffoldapplication"
	// ActionApplyApplication represents running the command `okctl apply application`
	ActionApplyApplication types.Action = "applyapplication"

	// ActionForwardPostgres represents running the command `okctl forward postgres`
	ActionForwardPostgres types.Action = "forwardpostgres"
	// ActionAttachPostgres represents running the command `okctl attach postgres`
	ActionAttachPostgres types.Action = "attachpostgres"

	// ActionShowCredentials represents running the command `okctl show credentials`
	ActionShowCredentials types.Action = "showcredentials"
	// ActionUpgrade represents running the command `okctl upgrade`
	ActionUpgrade types.Action = "upgradecluster"
	// ActionVenv represents running the command `okctl venv`
	ActionVenv types.Action = "venvcluster"
	// ActionVersion represents running the command `okctl version`
	ActionVersion types.Action = "version"

	// ActionMaintenanceStateAcquireLock represents running the command `okctl maintenance state-acquire-lock
	ActionMaintenanceStateAcquireLock types.Action = "stateacquirelock"
	// ActionMaintenanceStateReleaseLock represents running the command `okctl maintenance state-release-lock
	ActionMaintenanceStateReleaseLock types.Action = "statereleaselock"
	// ActionMaintenanceStateDownload represents running the command `okctl maintenance state-download
	ActionMaintenanceStateDownload types.Action = "statedownload"
	// ActionMaintenanceStateUpload represents running the command `okctl maintenance state-upload
	ActionMaintenanceStateUpload types.Action = "stateupload"
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
		ActionScaffoldCluster, ActionApplyCluster, ActionDeleteCluster,
		ActionScaffoldApplication, ActionApplyApplication,
		ActionForwardPostgres, ActionAttachPostgres,
		ActionShowCredentials, ActionUpgrade, ActionVenv, ActionVersion,
		ActionMaintenanceStateAcquireLock, ActionMaintenanceStateReleaseLock,
		ActionMaintenanceStateDownload, ActionMaintenanceStateUpload,
	},
	Labels: []string{LabelPhaseKey},
}
