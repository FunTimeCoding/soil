package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"testing"
)

func TestClaimingTwiceKeepsOneClaim(t *testing.T) {
	s := newStore(t, notifier.New())
	_, e := s.ClaimItem("jira-ABC-1", "alfa@host.example")
	assert.FatalOnError(t, e)
	_, e = s.ClaimItem("jira-ABC-1", "alfa@host.example")
	assert.FatalOnError(t, e)
	v, e := s.Claims()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
}

func TestTwoOwnersClaimSameItem(t *testing.T) {
	s := newStore(t, notifier.New())
	_, e := s.ClaimItem("jira-ABC-1", "alfa@host.example")
	assert.FatalOnError(t, e)
	_, e = s.ClaimItem("jira-ABC-1", "bravo@host.example")
	assert.FatalOnError(t, e)
	v, e := s.Claims()
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
}

func TestReleaseRemovesOnlyOwnersClaim(t *testing.T) {
	s := newStore(t, notifier.New())
	_, e := s.ClaimItem("jira-ABC-1", "alfa@host.example")
	assert.FatalOnError(t, e)
	_, e = s.ClaimItem("jira-ABC-1", "bravo@host.example")
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, s.Release("jira-ABC-1", "alfa@host.example"))
	v, e := s.Claims()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "bravo@host.example", v[0].Owner)
}
