//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/service_tester"
	"testing"
)

func TestEmbedCreatesEmbeddings(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	result, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, float64(result.Documents))
	assert.Greater(t, 0, float64(result.Chunks))
	status := s.Service.MustStatus()
	assert.Greater(t, 0, float64(status.TotalEmbeddings))
	assert.Integer(t, 0, status.PendingEmbeddings)
}

func TestEmbedIdempotent(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	first, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	second, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, second.Documents)
	assert.Integer(t, 0, second.Chunks)
	status := s.Service.MustStatus()
	assert.Integer(t, first.Chunks, status.TotalEmbeddings)
}

func TestEmbedAfterPush(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"embed-me.md",
			"# Embed Test\n\nThis document should get embedded.\n",
			nil,
		),
	)
	status := s.Service.MustStatus()
	assert.Greater(t, 0, float64(status.TotalEmbeddings))
}
