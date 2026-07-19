package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"testing"
)

func TestAlertmanagerAlertAdvancedOption(t *testing.T) {
	assert.NotNil(t, advanced_option.New())
}
