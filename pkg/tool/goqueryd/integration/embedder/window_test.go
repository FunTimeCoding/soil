//go:build local

package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/strings/join"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration/embedder_tester"
	"strings"
	"testing"
)

func TestEveryManualChunkFitsEmbedder(t *testing.T) {
	embedder_tester.Embedded(
		t,
		embedder_tester.New(t),
		"manual.txt",
		fixture.Read(systemConstant.LimitPath, "manual.txt"),
	)
}

func TestOllamaClientRefusesDenseChunk(t *testing.T) {
	embedder_tester.Refused(
		t,
		embedder_tester.NewOllama(t),
		"dense.md",
		fixture.Read(systemConstant.LimitPath, "dense.md"),
	)
}

func TestOllamaClientRefusesGoDeclaration(t *testing.T) {
	embedder_tester.Refused(
		t,
		embedder_tester.NewOllama(t),
		"table.go",
		join.Empty(
			"package alfa\n\nvar table = []int{\n",
			strings.Repeat("\t1, 2, 3, 4, 5, 6, 7, 8, 9, 0,\n", 110),
			"}\n",
		),
	)
}
