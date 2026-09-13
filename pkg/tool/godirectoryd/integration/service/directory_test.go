//go:build ci

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap(directory.NewEnvironment())
	os.Exit(m.Run())
}

func stack() *service.Service {
	return service.New(directory.NewEnvironment())
}

func TestUserLifecycle(t *testing.T) {
	s := stack()
	created, e := s.CreateUser(
		"lifecycle",
		"Lifecycle Person",
		"Person",
		"lifecycle@example.test",
		"lifecyclepassword",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "lifecycle", created.Account)
	assert.String(t, "lifecycle@example.test", created.Mail)

	if created.Unique == "" {
		t.Errorf("expected a unique identifier")
	}

	modified, f := s.ModifyUser("lifecycle", "", "", "moved@example.test")
	assert.FatalOnError(t, f)
	assert.String(t, "moved@example.test", modified.Mail)
	assert.String(t, created.Unique, modified.Unique)
	assert.FatalOnError(t, s.DeleteUser("lifecycle"))
	_, g := s.User("lifecycle")
	assert.True(t, not_found.Is(g))
}

func TestCreateUserDuplicateConflicts(t *testing.T) {
	s := stack()
	_, e := s.CreateUser("twice", "Twice", "Person", "", "")
	assert.FatalOnError(t, e)

	defer func() { assert.FatalOnError(t, s.DeleteUser("twice")) }()
	_, f := s.CreateUser("twice", "Twice", "Person", "", "")
	assert.True(t, conflict.Is(f))
}

func TestSetPasswordThenAuthenticate(t *testing.T) {
	s := stack()
	_, e := s.CreateUser("rotate", "Rotate", "Person", "", "firstpassword")
	assert.FatalOnError(t, e)

	defer func() { assert.FatalOnError(t, s.DeleteUser("rotate")) }()
	assert.FatalOnError(t, s.SetPassword("rotate", "secondpassword"))
	entry, f := directory.NewEnvironment().Authenticate(
		"rotate",
		"secondpassword",
	)
	assert.FatalOnError(t, f)
	assert.String(t, "rotate", entry.Account)
}

func TestGroupLifecycle(t *testing.T) {
	s := stack()
	created, e := s.CreateGroup("editors")
	assert.FatalOnError(t, e)
	assert.String(t, "editors", created.Name)
	assert.Greater(t, 4999, created.Number)

	defer func() { assert.FatalOnError(t, s.DeleteGroup("editors")) }()
	_, f := s.CreateUser("joiner", "Joiner", "Person", "", "")
	assert.FatalOnError(t, f)

	defer func() { assert.FatalOnError(t, s.DeleteUser("joiner")) }()
	joined, g := s.AddMember("editors", "joiner")
	assert.FatalOnError(t, g)
	assert.Count(t, 1, joined.Member)
	assert.String(t, "joiner", joined.Member[0])
	_, h := s.AddMember("editors", "joiner")
	assert.True(t, conflict.Is(h))
	left, j := s.RemoveMember("editors", "joiner")
	assert.FatalOnError(t, j)
	assert.Count(t, 0, left.Member)
}

func TestAddMemberUnknownUser(t *testing.T) {
	s := stack()
	_, e := s.CreateGroup("empty")
	assert.FatalOnError(t, e)

	defer func() { assert.FatalOnError(t, s.DeleteGroup("empty")) }()
	_, f := s.AddMember("empty", "ghost")
	assert.True(t, not_found.Is(f))
}

func TestGroupNumberIncrements(t *testing.T) {
	s := stack()
	first, e := s.CreateGroup("alpha")
	assert.FatalOnError(t, e)

	defer func() { assert.FatalOnError(t, s.DeleteGroup("alpha")) }()
	second, f := s.CreateGroup("beta")
	assert.FatalOnError(t, f)

	defer func() { assert.FatalOnError(t, s.DeleteGroup("beta")) }()
	assert.Integer(t, first.Number+1, second.Number)
}
