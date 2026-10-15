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

	Command          = "terraform"
	Init             = "init"
	NotationArgument = "-json"
	UpgradeArgument  = "-upgrade"
	ForceUnlock      = "force-unlock"
	ForceArgument    = "-force"
	LockContext      = "terraform_lock"

	TerraformPathEnvironment   = "TERRAFORM_PATH"
	DownstreamEnvironment      = "DOWNSTREAM_HOST"
	DownstreamTokenEnvironment = "DOWNSTREAM_TOKEN" // #nosec G101 not a hardcoded secret
	StateNamespaceEnvironment  = "TERRAFORM_STATE_NAMESPACE"
	StateLeaseEnvironment      = "TERRAFORM_STATE_LEASE"

	StateNamespace = "terraform"
	StateLease     = "lock-tfstate-default-state"
	LockAnnotation = "app.terraform.io/lock-info"
)

const (
	FixtureHostname        = "runner-6d4b8c9f7-abcde"
	FixtureLock            = "11111111-2222-4333-8444-666666666666"
	FixtureGonePod         = "runner-5c7d9e8f1-fghij"
	FixtureLaptop          = "user@workstation"
	FixtureLongOwn         = "a-deployment-name-reaching-the-hostname-cutoff-0987654321-zyxwv"
	FixtureLongOther       = "a-deployment-name-reaching-the-hostname-cutoff-1234567890-abcde"
	FixtureGoneHolder      = "root@runner-5c7d9e8f1-fghij"
	FixtureOwnHolder       = "root@runner-6d4b8c9f7-abcde"
	FixtureOtherHolder     = "root@service-alfa-7b6c5d4e3-klmno"
	FixtureLongOtherHolder = "root@a-deployment-name-reaching-the-hostname-cutoff-1234567890-abcde"
)
