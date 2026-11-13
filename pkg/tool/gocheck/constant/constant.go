package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gocheck",
	"Every lint stage in one run: golint, goaudit, gofix and goanalyze",
	"gocheck [flags]",
)

const (
	Reach       = "reach"
	Fresh       = "fresh"
	OwnSession  = "own"
	LoadFailure = "load_failure"
	CleanMarks  = "gocheck"

	RewrittenSuffix = " - on a line the fixes rewrite"
	NothingChanged  = "gocheck: nothing changed"
	StaleNotice     = "gocheck: gosourced started before %s changed - restart it, or pass --local\n"
)
