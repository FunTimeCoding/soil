package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"gosaltd",
	"Salt highstate runner",
	"gosaltd",
).WithInstructions(
	"Salt highstate runner - trigger highstate on all minions and check results. Call sync to pull latest code. Use trigger with update and synchronous flags for single-call deploy. Runs execute asynchronously by default; use runs and run tools to poll for completion and read output.",
)

const (
	Trigger     = "trigger"
	Sync        = "sync"
	Runs        = "runs"
	Run         = "run"
	Minions     = "minions"
	Keys        = "keys"
	Target      = "target"
	Update      = "update"
	Synchronous = "synchronous"
	Limit       = "limit"
	Status      = "status"
	Identifier  = "id"

	SaltDirectoryEnvironment = "SALT_DIRECTORY"

	SaltPeer = "salt"

	ConnectRetryInterval = 30 * time.Second
	KeySyncInterval      = 30 * time.Second
)
