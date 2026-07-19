package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/lower"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"testing"
)

func TestSyncAddsNewFiles(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{
				Name:        lower.Alfa,
				Path:        "alfa.md",
				ContentHash: "hash-a",
				Content:     "a",
			},
			{
				Name:        lower.Bravo,
				Path:        "bravo.md",
				ContentHash: "hash-b",
				Content:     "b",
			},
		},
	)
	seeds := s.Service.Seeds()
	assert.Count(t, 2, seeds)
	assert.String(t, "alfa", seeds[0].Name)
	assert.String(t, "bravo", seeds[1].Name)
	assert.True(t, s.Notifier.Notified > 0)
}

func TestSyncRemovesDeletedFiles(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{
				Name:        lower.Alfa,
				Path:        "alfa.md",
				ContentHash: "hash-a",
				Content:     "a",
			},
			{
				Name:        lower.Bravo,
				Path:        "bravo.md",
				ContentHash: "hash-b",
				Content:     "b",
			},
		},
	)
	s.Notifier.Reset()
	s.Service.Sync(
		[]service.DiscoveredFile{
			{
				Name:        lower.Alfa,
				Path:        "alfa.md",
				ContentHash: "hash-a",
				Content:     "a",
			},
		},
	)
	seeds := s.Service.Seeds()
	assert.Count(t, 1, seeds)
	assert.String(t, "alfa", seeds[0].Name)
	assert.True(t, s.Notifier.Notified > 0)
}

func TestSyncUpdatesContent(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{
				Name:        lower.Alfa,
				Path:        "alfa.md",
				ContentHash: "hash-1",
				Content:     "old",
			},
		},
	)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{
				Name:        lower.Alfa,
				Path:        "alfa.md",
				ContentHash: "hash-2",
				Content:     "new",
			},
		},
	)
	seeds := s.Service.Seeds()
	assert.String(t, "new", seeds[0].Content)
}

func TestSyncPreservesPositionOnUpdate(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{
				Name:        lower.Alfa,
				Path:        "alfa.md",
				ContentHash: "hash-a",
				Content:     "a",
			},
			{
				Name:        lower.Bravo,
				Path:        "bravo.md",
				ContentHash: "hash-b",
				Content:     "b",
			},
		},
	)
	s.Service.SetPosition(s.Service.Seeds()[1].Identifier, 1)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{
				Name:        lower.Alfa,
				Path:        "alfa.md",
				ContentHash: "hash-a2",
				Content:     "a2",
			},
			{
				Name:        lower.Bravo,
				Path:        "bravo.md",
				ContentHash: "hash-b",
				Content:     "b",
			},
		},
	)
	seeds := s.Service.Seeds()
	assert.String(t, "bravo", seeds[0].Name)
	assert.String(t, "alfa", seeds[1].Name)
}
