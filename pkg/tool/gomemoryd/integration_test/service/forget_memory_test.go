//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/service_tester"
	"testing"
)

func TestServiceForgetMemory(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(
		"to forget",
		constant.FixtureContent,
		"desc",
		"feedback",
		"test",
		nil,
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, o.Service.ForgetMemory(m.Identifier, "test"))
	assert.Count(t, 1, o.Indexer.Deleted)
	assert.String(t, "memory/1", o.Indexer.Deleted[0])
	active, e := o.Service.ListMemories(
		"",
		"",
		true,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, active)
}
