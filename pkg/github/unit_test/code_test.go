package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/code"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/google/go-github/v90/github"
	"testing"
)

func TestCode(t *testing.T) {
	r := code.New(
		&github.CodeResult{
			SHA:  new(constant.UpperAlfa),
			Name: new(constant.UpperBravo),
			Path: new(constant.UpperCharlie),
		},
	)
	r.Raw = nil
	assert.Any(t, &code.Code{Hash: "Alfa", Name: "Bravo", Path: "Charlie"}, r)
}
