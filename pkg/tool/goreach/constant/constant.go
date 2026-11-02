package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goreach",
	"Directories a session's changes reach: its packages and their importers",
	"goreach [flags] [session]",
)

const (
	Transcript = "transcript"
	Documents  = "documents"
)
