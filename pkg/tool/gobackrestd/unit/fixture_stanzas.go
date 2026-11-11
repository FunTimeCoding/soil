package unit

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/backrest/response"
	"github.com/funtimecoding/soil/pkg/notation"
)

func fixtureStanzas(name string) []*response.Stanza {
	var result []*response.Stanza
	notation.MustDecode(fixture.Read("backrest", name), &result, false)

	return result
}
