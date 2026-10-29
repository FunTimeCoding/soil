package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/service_tester"
	"testing"
)

func TestServiceRejectsAClaimWithoutAnItem(t *testing.T) {
	_, e := service_tester.New(t).ClaimItem("", "alfa@host.example")
	assert.True(t, validation.Is(e))
}

func TestServiceRejectsAReleaseWithoutAnOwner(t *testing.T) {
	assert.True(
		t,
		validation.Is(service_tester.New(t).Release("jira-ABC-1", "")),
	)
}

func TestServiceClaimsThroughItsOwnDatabase(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.ClaimItem("jira-ABC-1", "alfa@host.example")
	assert.FatalOnError(t, e)
	v, f := s.Claims()
	assert.FatalOnError(t, f)
	assert.Count(t, 1, v)
}
