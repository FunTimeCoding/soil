package go_mod

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestExtractReplaces(t *testing.T) {
	assert.String(
		t,
		"replace (\n\tk8s.io/api => k8s.io/api v0.31.0\n)",
		ExtractReplaces(
			`module github.com/example/project/v2

go 1.22.0

require (
	golang.org/x/crypto v0.31.0
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f
)

replace (
	k8s.io/api => k8s.io/api v0.31.0
)`,
		),
	)
}
