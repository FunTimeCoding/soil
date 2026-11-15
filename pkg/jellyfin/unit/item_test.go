package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/jellyfin/unit/jellyfin_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestItemIsFoundByItsIdentifier(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"Items":[{"Id":"a1b2c3d4e5f60718293a4b5c6d7e8f90","Name":"bravo"}],"TotalRecordCount":1}`,
		),
	)
	i, e := jellyfin_tester.New(t, s.URL).Item(
		"A1B2C3D4-E5F6-0718-293A-4B5C6D7E8F90",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "bravo", i.Name)
}

func TestUnknownItemIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"Items":[],"TotalRecordCount":0}`),
	)
	_, e := jellyfin_tester.New(t, s.URL).Item(
		"0123456789abcdef0123456789abcdef",
	)
	assert.True(t, not_found.Is(e))
	assert.String(
		t,
		"item not found: 0123456789abcdef0123456789abcdef",
		e.Error(),
	)
}

func TestMalformedIdentifierIsNotFoundDespiteUnrelatedItems(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"Items":[{"Id":"0f0e0d0c0b0a09080706050403020100","Name":"delta"}],"TotalRecordCount":1}`,
		),
	)
	_, e := jellyfin_tester.New(t, s.URL).Item("alfa")
	assert.True(t, not_found.Is(e))
}
