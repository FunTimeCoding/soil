package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosproutd",
	"Local seed priority tracker",
	"gosproutd",
).WithInstructions(
	"Local seed priority tracker. Scans a seed directory for markdown files, tracks them with priority ordering. Use list_seeds to see priorities, reorder_seed to change them.",
)

const (
	SeedDirectoryEnvironment = "SEED_DIRECTORY"

	ListSeeds   = "list_seeds"
	ReorderSeed = "reorder_seed"

	DashboardTitle = "Dashboard"
	DashboardPath  = "/"

	Seeds = "seeds"
)
