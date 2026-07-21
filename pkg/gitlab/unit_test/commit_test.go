package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestCommit(t *testing.T) {
	c := commit.New(&gitlab.Commit{})
	c.Raw = nil
	assert.Exported(t, &commit.Commit{}, c)
}
