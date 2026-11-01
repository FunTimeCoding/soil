//go:build local

package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/strings/join"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
	"testing"
)

func TestEveryManualChunkFitsTheEmbedder(t *testing.T) {
	assertEmbedded(
		t,
		newEmbedder(t),
		"manual.txt",
		fixture.Read(systemConstant.LimitPath, "manual.txt"),
	)
}

func TestTheOllamaClientRefusesTheDenseChunk(t *testing.T) {
	assertRefused(
		t,
		newOllamaEmbedder(t),
		"dense.md",
		fixture.Read(systemConstant.LimitPath, "dense.md"),
	)
}

func TestTheOllamaClientRefusesTheGoDeclaration(t *testing.T) {
	assertRefused(
		t,
		newOllamaEmbedder(t),
		"table.go",
		join.Empty(
			"package alfa\n\nvar table = []int{\n",
			strings.Repeat("\t1, 2, 3, 4, 5, 6, 7, 8, 9, 0,\n", 110),
			"}\n",
		),
	)
}
