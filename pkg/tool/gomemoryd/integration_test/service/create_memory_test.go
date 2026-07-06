//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/service_tester"
	"testing"
)

func TestServiceCreateMemory(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(
		"retry policy",
		"Retry failed requests with exponential backoff.",
		"retry with backoff",
		"",
		"test",
		nil,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "retry policy", m.Name)
	assert.String(t, "feedback", m.Type)
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.String(t, "memory/1", o.Indexer.Pushed[0].Name)
}

func TestServiceCreateMemoryWithExplicitType(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(
		"deploy target",
		"Production deployments use blue-green strategy.",
		"blue-green deploy pattern",
		"user",
		"test",
		nil,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "user", m.Type)
}
