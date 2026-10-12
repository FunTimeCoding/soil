package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestDeleteRegistryPathRemovesThePathAndItsCache(t *testing.T) {
	c := mock_client.New()
	c.SeedRegistryRepository(7, 1, "group/app")
	c.SeedRegistryRepository(7, 2, "group/app/cache")
	c.SeedRegistryRepository(7, 3, "group/other")
	s := service_tester.Service(c)
	deleted, e := s.DeleteRegistryPath(7, "group/app")
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"group/app", "group/app/cache"}, deleted)
	remaining, f := s.RegistryRepositories(7, false)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, remaining)
	assert.String(t, "group/other", remaining[0].Path)
}

func TestDeleteRegistryPathWithoutACacheDeletesWhatExists(t *testing.T) {
	c := mock_client.New()
	c.SeedRegistryRepository(7, 1, "group/app")
	deleted, e := service_tester.Service(c).DeleteRegistryPath(7, "group/app")
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"group/app"}, deleted)
}

func TestDeleteRegistryPathRefusesWhenNothingMatches(t *testing.T) {
	c := mock_client.New()
	c.SeedRegistryRepository(7, 1, "group/app")
	s := service_tester.Service(c)
	_, e := s.DeleteRegistryPath(7, "group/missing")
	assert.True(t, not_found.Is(e))
	assert.String(t, "registry repository not found: group/missing", e.Error())
	remaining, f := s.RegistryRepositories(7, false)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, remaining)
}

func TestDeleteRegistryPathStaysInsideItsProject(t *testing.T) {
	c := mock_client.New()
	c.SeedRegistryRepository(7, 1, "group/app")
	c.SeedRegistryRepository(8, 2, "group/app")
	s := service_tester.Service(c)
	_, e := s.DeleteRegistryPath(7, "group/app")
	assert.FatalOnError(t, e)
	other, f := s.RegistryRepositories(8, false)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, other)
}
