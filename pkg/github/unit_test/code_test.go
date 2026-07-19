package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/code"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
)

func TestCode(t *testing.T) {
	r := code.New(
		&github.CodeResult{
			SHA:  new(upper.Alfa),
			Name: new(upper.Bravo),
			Path: new(upper.Charlie),
		},
	)
	r.Raw = nil
	assert.Any(t, &code.Code{Hash: "Alfa", Name: "Bravo", Path: "Charlie"}, r)
}
