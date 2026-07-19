package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_processor"
	"testing"
)

func TestProcessor(t *testing.T) {
	assert.NotNil(t, alert_processor.New(nil, nil, nil, nil, nil, nil))
}
