package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestLokiClient(t *testing.T) {
	assert.NotNil(
		t,
		loki.New(upper.Alfa, upper.Bravo, upper.Charlie, false),
	)
}
