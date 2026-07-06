package outdated

import (
	"github.com/funtimecoding/soil/pkg/system/macos/brew"
	"github.com/funtimecoding/soil/pkg/system/macos/brew/formula"
)

func collect() []*formula.Formula {
	c := brew.New()
	r := c.Outdated()
	var result []*formula.Formula

	for _, e := range r.Formulae {
		result = append(result, formula.New(e))
	}

	return result
}
