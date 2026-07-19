package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"testing"
)

func TestReplaceReplaces(t *testing.T) {
	assert.String(
		t,
		"replace (\nb\n)\n",
		go_mod.ReplaceReplaces(
			"replace (\na\n)\n",
			"replace (\nb\n)\n",
		),
	)
}
