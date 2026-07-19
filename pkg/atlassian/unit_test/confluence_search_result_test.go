package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"
	"testing"
)

func TestResult(t *testing.T) {
	assert.NotNil(t, search_result.New(response.NewResult()))
}
