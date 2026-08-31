package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goterraformd",
	"Terraform apply runner",
	"goterraformd",
).WithInstructions(
	"Terraform apply runner - trigger init + apply and check results. Call sync to pull latest code. Use trigger with update and synchronous flags for single-call deploy. Runs execute asynchronously by default; use runs and run tools to poll for completion and read output.",
)

const (
	Trigger     = "trigger"
	Sync        = "sync"
	Runs        = "runs"
	Run         = "run"
	Target      = "target"
	Update      = "update"
	Synchronous = "synchronous"
	Limit       = "limit"
	Status      = "status"
	Identifier  = "id"

	Command = "terraform"

	TerraformPathEnvironment  = "TERRAFORM_PATH"
	DownstreamEnvironment     = "DOWNSTREAM_HOSTS"
	StateNamespaceEnvironment = "TERRAFORM_STATE_NAMESPACE"
	StateLeaseEnvironment     = "TERRAFORM_STATE_LEASE"

	StateNamespace = "terraform"
	StateLease     = "lock-tfstate-default-state"
	LockAnnotation = "app.terraform.io/lock-info"
)
