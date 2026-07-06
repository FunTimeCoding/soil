package space

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"testing"
)

func TestSpace(t *testing.T) {
	assert.NotNil(t, New(response.NewSpace()))
}
