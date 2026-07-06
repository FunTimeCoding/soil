package lint

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"strings"
	"testing"
)

func TestVariableGroupingConsecutive(t *testing.T) {
	r := VariableGrouping(
		"test.go",
		strings.NewReader("package main\n\nvar a = 1\nvar b = 2\n"),
	)
	assert.Boolean(t, true, r.HasConcerns())
	assert.String(
		t,
		"package main\n\nvar (\n\ta = 1\n\tb = 2\n)\n",
		r.Fixed,
	)
}

func TestVariableGroupingThree(t *testing.T) {
	r := VariableGrouping(
		"test.go",
		strings.NewReader("package main\n\nvar a = 1\nvar b = 2\nvar c = 3\n"),
	)
	assert.String(
		t,
		"package main\n\nvar (\n\ta = 1\n\tb = 2\n\tc = 3\n)\n",
		r.Fixed,
	)
}

func TestVariableGroupingSingleNotGrouped(t *testing.T) {
	r := VariableGrouping(
		"test.go",
		strings.NewReader("package main\n\nvar a = 1\n"),
	)
	assert.Boolean(t, false, r.HasConcerns())
}

func TestVariableGroupingMultiLineSkipped(t *testing.T) {
	r := VariableGrouping(
		"test.go",
		strings.NewReader("package main\n\nvar a = []string{\n\t\"x\",\n}\nvar b = 2\n"),
	)
	assert.Boolean(t, false, r.HasConcerns())
}

func TestVariableGroupingSeparatedByBlank(t *testing.T) {
	r := VariableGrouping(
		"test.go",
		strings.NewReader("package main\n\nvar a = 1\n\nvar b = 2\n"),
	)
	assert.Boolean(t, false, r.HasConcerns())
}

func TestVariableGroupingBlankIdentifier(t *testing.T) {
	r := VariableGrouping(
		"test.go",
		strings.NewReader("package main\n\nvar _ = 1\nvar _ = 2\n"),
	)
	assert.String(
		t,
		"package main\n\nvar (\n\t_ = 1\n\t_ = 2\n)\n",
		r.Fixed,
	)
}
