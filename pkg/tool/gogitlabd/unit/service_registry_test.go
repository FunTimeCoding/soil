package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestDeleteRegistryPathRemovesThePathAndItsCache(t *testing.T) {
	s, clients := service_tester.New(t, constant.FixtureInstance)
	c := clients[constant.FixtureInstance]
	c.SeedRegistryRepository(7, 1, "group/app")
	c.SeedRegistryRepository(7, 2, "group/app/cache")
	c.SeedRegistryRepository(7, 3, "group/other")
	deleted, e := s.DeleteRegistryPath(constant.FixtureInstance, 7, "group/app")
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"group/app", "group/app/cache"}, deleted)
	remaining, f := s.RegistryRepositories(constant.FixtureInstance, 7, false)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, remaining)
	assert.String(t, "group/other", remaining[0].Path)
}

func TestDeleteRegistryPathWithoutACacheDeletesWhatExists(t *testing.T) {
	s, clients := service_tester.New(t, constant.FixtureInstance)
	clients[constant.FixtureInstance].SeedRegistryRepository(7, 1, "group/app")
	deleted, e := s.DeleteRegistryPath(constant.FixtureInstance, 7, "group/app")
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"group/app"}, deleted)
}

func TestDeleteRegistryPathRefusesWhenNothingMatches(t *testing.T) {
	s, clients := service_tester.New(t, constant.FixtureInstance)
	clients[constant.FixtureInstance].SeedRegistryRepository(7, 1, "group/app")
	_, e := s.DeleteRegistryPath(constant.FixtureInstance, 7, "group/missing")
	assert.True(t, not_found.Is(e))
	assert.String(t, "registry repository not found: group/missing", e.Error())
	remaining, f := s.RegistryRepositories(constant.FixtureInstance, 7, false)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, remaining)
}

func TestDeleteRegistryPathStaysInsideItsProject(t *testing.T) {
	s, clients := service_tester.New(t, constant.FixtureInstance)
	c := clients[constant.FixtureInstance]
	c.SeedRegistryRepository(7, 1, "group/app")
	c.SeedRegistryRepository(8, 2, "group/app")
	_, e := s.DeleteRegistryPath(constant.FixtureInstance, 7, "group/app")
	assert.FatalOnError(t, e)
	other, f := s.RegistryRepositories(constant.FixtureInstance, 8, false)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, other)
}
