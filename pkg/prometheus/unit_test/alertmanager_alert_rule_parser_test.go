package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/rule_parser"
	"testing"
)

func TestParser(t *testing.T) {
	assert.NotNil(t, rule_parser.New(nil))
}
