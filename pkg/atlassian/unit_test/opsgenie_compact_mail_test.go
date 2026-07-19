package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/compact"
	"testing"
)

func TestMail(t *testing.T) {
	assert.String(t, "alfa", compact.Mail("alfa@bravo"))
}
