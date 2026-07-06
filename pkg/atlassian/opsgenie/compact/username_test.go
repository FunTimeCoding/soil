package compact

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestUsername(t *testing.T) {
	// noinspection SpellCheckingInspection
	assert.String(t, "abravo", Username("alfa.bravo@charlie"))
}
