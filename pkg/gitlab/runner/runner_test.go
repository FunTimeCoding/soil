package runner

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestRunner(t *testing.T) {
	r := New(&gitlab.Runner{})
	r.RawList = nil
	assert.Any(t, &Runner{}, r)
}
