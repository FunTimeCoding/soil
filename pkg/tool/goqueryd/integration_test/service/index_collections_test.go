//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/service_tester"
	"testing"
)

func TestIndexCollections(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	results, e := s.Service.IndexCollections("")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, results)
	assert.String(t, "test", results[0].Collection)
	assert.Integer(t, 0, results[0].Indexed)
	assert.Integer(t, 5, results[0].Unchanged)
}

func TestIndexCollectionsFilter(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	results, e := s.Service.IndexCollections("nonexistent")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, results)
}
