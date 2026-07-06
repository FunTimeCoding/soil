package enrichment

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestEnrichment(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa, upper.Bravo, upper.Charlie))
}
