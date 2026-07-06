package job

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	assert.Count(t, 0, Deduplicate([]*gitlab.Job{}))
}
