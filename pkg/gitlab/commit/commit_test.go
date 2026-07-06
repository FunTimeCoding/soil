package commit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestCommit(t *testing.T) {
	c := New(&gitlab.Commit{})
	c.Raw = nil
	assert.Any(t, &Commit{}, c)
}
