package rule_list

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestList(t *testing.T) {
	assert.NotNil(t, New())
}
