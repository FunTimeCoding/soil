package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"testing"
)

func TestJoinBlockMergesGroupsSorted(t *testing.T) {
	assert.String(
		t,
		"package alfa\n\nimport (\n\t\"example.org/m/pkg/bravo\"\n\tcharlie \"example.org/m/pkg/charlie\"\n\t\"fmt\"\n)\n",
		string(
			imports.JoinBlock(
				[]byte(
					"package alfa\n\nimport (\n\t\"fmt\"\n\n\tcharlie \"example.org/m/pkg/charlie\"\n\t\"example.org/m/pkg/bravo\"\n)\n",
				),
			),
		),
	)
}

func TestJoinBlockLeavesASingleBlock(t *testing.T) {
	source := "package alfa\n\nimport (\n\t\"example.org/m/pkg/bravo\"\n\t\"fmt\"\n)\n"
	assert.String(t, source, string(imports.JoinBlock([]byte(source))))
}

func TestJoinBlockLeavesACommentedBlock(t *testing.T) {
	source := "package alfa\n\nimport (\n\t\"fmt\"\n\n\t// the bravo group\n\t\"example.org/m/pkg/bravo\"\n)\n"
	assert.String(t, source, string(imports.JoinBlock([]byte(source))))
}
