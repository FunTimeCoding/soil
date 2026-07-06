package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goansibled",
	"Ansible playbook runner",
	"goansibled",
).WithInstructions(
	"Ansible playbook runner - trigger runs and check results. Call sync to pull latest code. Use trigger with update and synchronous flags for single-call deploy. Runs execute asynchronously by default; use runs and run tools to poll for completion and read output.",
)

const (
	Playbooks   = "playbooks"
	Trigger     = "trigger"
	Sync        = "sync"
	Runs        = "runs"
	Run         = "run"
	Playbook    = "playbook"
	Update      = "update"
	Synchronous = "synchronous"
	Limit       = "limit"
	Status      = "status"
	Identifier  = "id"

	AnsiblePathEnvironment = "ANSIBLE_PATH"

	RecentRunsFailed = "failed to list recent runs"
	RunLookupFailed  = "failed to look up run"
)
