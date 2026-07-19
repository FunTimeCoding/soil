package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/field_changer/severity"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestSeverity(t *testing.T) {
	assert.NotNil(t, severity.New(upper.Alfa, upper.Bravo, upper.Charlie))
}
