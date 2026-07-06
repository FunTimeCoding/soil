package rule_parser

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestParser(t *testing.T) {
	assert.NotNil(t, New(nil))
}
