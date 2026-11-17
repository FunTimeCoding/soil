//go:build local

package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/strings/join"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration/embedder_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"strings"
	"testing"
)

func TestCounterAllowsWindowLessSpecials(t *testing.T) {
	assert.Integer(t, 2046, embedder_tester.NewCounter(t).TokenLimit())
}

func TestCounterSeesDenseChunkOverWindow(t *testing.T) {
	assert.Greater(
		t,
		2046,
		embedder_tester.NewCounter(t).Count(
			fixture.Read(systemConstant.LimitPath, "dense.md"),
		),
	)
}

func TestCounterSeesGoDeclarationOverWindow(t *testing.T) {
	assert.Greater(
		t,
		2046,
		embedder_tester.NewCounter(t).Count(
			join.Empty(
				"package alfa\n\nvar table = []int{\n",
				strings.Repeat("\t1, 2, 3, 4, 5, 6, 7, 8, 9, 0,\n", 110),
				"}\n",
			),
		),
	)
}

func TestCounterFitsEveryManualChunk(t *testing.T) {
	c := embedder_tester.NewCounter(t)

	for _, k := range chunk.Document(
		fixture.Read(systemConstant.LimitPath, "manual.txt"),
		"manual.txt",
		service_tester.New(t).Reranker,
	) {
		assert.LessEqual(t, 2046, c.Count(k.Text))
	}
}
