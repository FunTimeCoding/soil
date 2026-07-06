package alert_filter

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"testing"
)

func TestFilter(t *testing.T) {
	// TODO: Test cases
	o := advanced_option.New()
	assert.Any(
		t,
		[]*alert.Alert{},
		New(o).Run([]*alert.Alert{}),
	)
}
