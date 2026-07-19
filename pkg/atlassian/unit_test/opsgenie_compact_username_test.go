package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/compact"
	"testing"
)

func TestUsername(t *testing.T) {
	// noinspection SpellCheckingInspection
	assert.String(t, "abravo", compact.Username("alfa.bravo@charlie"))
}
