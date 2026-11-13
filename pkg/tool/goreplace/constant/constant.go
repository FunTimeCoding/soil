package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goreplace",
	"Apply search/replace blocks from stdin to one file, all or none",
	"goreplace",
)

const (
	Failed        = "%s: nothing written\n%s"
	BlockHeader   = "@@ block %d, line %d"
	Applied       = "%s: %d blocks applied"
	DryRunApplied = "%s: %d blocks would apply (dry run, nothing written)"
	RemovedPrefix = "-"
	AddedPrefix   = "+"
)
