package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"testing"
)

func TestSpace(t *testing.T) {
	assert.NotNil(t, space.New(response.NewSpace()))
}
