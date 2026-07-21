package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_filter"
	"testing"
)

func TestAlertFilter(t *testing.T) {
	// TODO: Test cases
	o := advanced_option.New()
	assert.Any(
		t,
		[]*alert.Alert{},
		alert_filter.New(o).Run([]*alert.Alert{}),
	)
}
