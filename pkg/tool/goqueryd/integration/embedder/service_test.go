//go:build local

package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestGoquerydEmbedsTheWholeManual(t *testing.T) {
	v := newService(t)
	assert.String(
		t,
		"",
		v.PushDocument(
			t.Context(),
			"manual",
			"manual.txt",
			fixture.Read(systemConstant.LimitPath, "manual.txt"),
			nil,
		),
	)
	assert.Integer(t, 3, v.MustStatus().TotalEmbeddings)
}

func TestGoquerydRefusesTheDenseChunkItself(t *testing.T) {
	v := newService(t)
	v.PushDocument(
		t.Context(),
		"dense",
		"dense.md",
		fixture.Read(systemConstant.LimitPath, "dense.md"),
		nil,
	)
	failed, e := v.Failed("")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, failed)
	assert.StringContains(t, "over the embedder's window", failed[0].Reason)
}
