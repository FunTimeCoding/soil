package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/option"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestOpsgenieAlertOptionAlert(t *testing.T) {
	assert.NotNil(
		t,
		option.New(
			nil,
			nil,
			upper.Alfa,
			nil,
			nil,
			nil,
			nil,
		),
	)
}
