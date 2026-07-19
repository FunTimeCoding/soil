package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/check/alert/option"
	"testing"
)

func TestCheckAlertOption(t *testing.T) {
	assert.NotNil(t, option.New())
}
