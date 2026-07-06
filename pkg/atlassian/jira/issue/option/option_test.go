package option

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestOption(t *testing.T) {
	assert.NotNil(
		t,
		New(
			upper.Alfa,
			upper.Bravo,
			[]string{},
			[]string{constant.Done},
			nil,
		),
	)
}
