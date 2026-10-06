package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/stamp/report"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"testing"
)

func TestACobraToolPrintsTheSameVersionBlockAsEveryTool(t *testing.T) {
	assert.String(t, stamp.New().Text(), cobraVersion(t, "--version"))
}

func TestACobraToolPrintsTheVersionAsNotation(t *testing.T) {
	assert.String(
		t,
		join.Empty(
			string(notation.MarshalIndent(report.New("gotest", stamp.New()))),
			"\n",
		),
		cobraVersion(t, "--version", "--notation"),
	)
}
