package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosource",
	"CLI for the gosourced Go source daemon",
	"gosource [command]",
)

const (
	ModuleFlag = "module"
	FromFlag   = "from"
	ToFlag     = "to"
	TextFlag   = "text"
)

const (
	CountsLabel   = "counts"
	UnparsedLabel = "unparsed"
	FromSide      = "from"
	ToSide        = "to"
	SourceArrow   = "<-"
	ModuleLabel   = "module"
	StagedLabel   = "staged"
	TextIndent    = "    "
)
