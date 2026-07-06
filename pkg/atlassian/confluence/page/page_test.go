package page

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestPage(t *testing.T) {
	assert.NotNil(t, New(response.NewPage(), upper.Alfa))
}
