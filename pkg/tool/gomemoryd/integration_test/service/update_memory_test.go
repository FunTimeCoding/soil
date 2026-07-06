//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/service_tester"
	"testing"
)

func TestServiceUpdateMemoryPreservesTags(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(
		"pace",
		"original content",
		"original desc",
		"feedback",
		"test",
		nil,
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		o.Service.AddTags(
			m.Identifier,
			[]string{"always", "go-conventions"},
		),
	)
	updated, e := o.Service.UpdateMemory(
		m.Identifier,
		"pace",
		"updated content",
		"updated desc",
		"test",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "updated content", updated.Content)
	assert.Count(t, 2, updated.Tags)
	assert.Count(t, 2, o.Indexer.Pushed)
}

func TestServiceUpdateMemoryNonexistentFails(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service.UpdateMemory(
		999,
		constant.FixtureName,
		constant.FixtureContent,
		"desc",
		"test",
	)
	assert.Error(t, e)
}
