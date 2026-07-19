package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"testing"
)

func TestList(t *testing.T) {
	assert.NotNil(t, rule_list.New())
}
