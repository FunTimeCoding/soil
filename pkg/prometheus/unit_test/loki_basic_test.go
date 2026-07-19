package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestLokiBasicClient(t *testing.T) {
	assert.NotNil(
		t,
		basic.New(upper.Alfa, upper.Bravo, upper.Charlie, false),
	)
}
