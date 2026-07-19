package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestRunner(t *testing.T) {
	r := runner.New(&gitlab.Runner{})
	r.RawList = nil
	assert.Any(t, &runner.Runner{}, r)
}
