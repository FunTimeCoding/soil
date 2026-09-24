package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gocrap",
	"CRAP report: cyclomatic complexity against test coverage per function",
	"gocrap <command> [flags] [pattern...]",
)

const (
	ScoreUsage     = "score [pattern...]"
	AttributeUsage = "attribute [pattern...]"
)
