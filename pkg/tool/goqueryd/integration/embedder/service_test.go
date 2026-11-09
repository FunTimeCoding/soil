//go:build local

package embedder

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestGoquerydEmbedsWholeManual(t *testing.T) {
	v := newService(t)
	assert.String(
		t,
		"",
		v.MustPushDocument(
			t.Context(),
			"manual",
			"manual.txt",
			fixture.Read(systemConstant.LimitPath, "manual.txt"),
			nil,
		),
	)
	assert.Integer(t, 3, v.MustStatus().TotalEmbeddings)
}

func TestGoquerydRefusesDenseChunkItself(t *testing.T) {
	v := newService(t)
	v.MustPushDocument(
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
