package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

type exportedFixture struct {
	Value  string
	hidden string
}

func TestExported(t *testing.T) {
	assert.Exported(
		t,
		&exportedFixture{Value: "a"},
		&exportedFixture{Value: "a"},
	)
}

func TestExportedIgnoresPrivateFields(t *testing.T) {
	assert.Exported(
		t,
		&exportedFixture{Value: "a", hidden: "b"},
		&exportedFixture{Value: "a", hidden: "c"},
	)
}
