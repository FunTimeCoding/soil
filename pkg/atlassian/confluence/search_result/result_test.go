package search_result

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"testing"
)

func TestResult(t *testing.T) {
	assert.NotNil(t, New(response.NewResult()))
}
