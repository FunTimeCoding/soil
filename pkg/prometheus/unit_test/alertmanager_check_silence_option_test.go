package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/option"
	"testing"
)

func TestAlertmanagerCheckSilenceOption(t *testing.T) {
	assert.NotNil(t, option.New())
}
