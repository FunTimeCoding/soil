package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestIssue(t *testing.T) {
	i := issue.New(&gitlab.Issue{})
	i.Raw = nil
	assert.Any(t, &issue.Issue{}, i)
}
