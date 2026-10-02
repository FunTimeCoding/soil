package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goflow",
	"Markdown prose reflow",
	"goflow [--width n] [--check] <file> [file ...]",
)

const (
	Width        = 80
	Mode         = 0o644
	Delimiter    = "---"
	FixtureWidth = 40
	Pipe         = "|"
)
