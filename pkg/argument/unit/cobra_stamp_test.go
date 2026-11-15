package unit

import (
	"github.com/funtimecoding/soil/pkg/argument/unit/argument_tester"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/stamp/report"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"testing"
)

func TestCobraToolPrintsSameVersionBlockAsEveryTool(t *testing.T) {
	assert.String(
		t,
		stamp.New().Text(),
		argument_tester.Version(t, "--version"),
	)
}

func TestCobraToolPrintsVersionAsNotation(t *testing.T) {
	assert.String(
		t,
		join.Empty(
			string(notation.MarshalIndent(report.New("gotest", stamp.New()))),
			"\n",
		),
		argument_tester.Version(t, "--version", "--notation"),
	)
}
