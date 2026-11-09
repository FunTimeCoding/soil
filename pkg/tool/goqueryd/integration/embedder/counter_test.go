//go:build local

package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/strings/join"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"strings"
	"testing"
)

func TestCounterAllowsWindowLessSpecials(t *testing.T) {
	assert.Integer(t, 2046, newCounter(t).Allowance())
}

func TestCounterSeesDenseChunkOverWindow(t *testing.T) {
	assert.Greater(
		t,
		2046,
		newCounter(t).Count(fixture.Read(systemConstant.LimitPath, "dense.md")),
	)
}

func TestCounterSeesGoDeclarationOverWindow(t *testing.T) {
	assert.Greater(
		t,
		2046,
		newCounter(t).Count(
			join.Empty(
				"package alfa\n\nvar table = []int{\n",
				strings.Repeat("\t1, 2, 3, 4, 5, 6, 7, 8, 9, 0,\n", 110),
				"}\n",
			),
		),
	)
}

func TestCounterFitsEveryManualChunk(t *testing.T) {
	c := newCounter(t)

	for _, k := range chunk.Document(
		fixture.Read(systemConstant.LimitPath, "manual.txt"),
		"manual.txt",
		newReranker(t),
	) {
		assert.LessEqual(t, 2046, c.Count(k.Text))
	}
}
