package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, option.New())
}
