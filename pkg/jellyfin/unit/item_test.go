package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"testing"
)

func TestItemIsFoundByItsIdentifier(t *testing.T) {
	i, e := items(
		t,
		`{"Items":[{"Id":"a1b2c3d4e5f60718293a4b5c6d7e8f90","Name":"bravo"}],"TotalRecordCount":1}`,
	).Item("A1B2C3D4-E5F6-0718-293A-4B5C6D7E8F90")
	assert.FatalOnError(t, e)
	assert.String(t, "bravo", i.Name)
}

func TestAnUnknownItemIsNotFound(t *testing.T) {
	_, e := items(t, `{"Items":[],"TotalRecordCount":0}`).Item(
		"0123456789abcdef0123456789abcdef",
	)
	assert.True(t, not_found.Is(e))
	assert.String(
		t,
		"item not found: 0123456789abcdef0123456789abcdef",
		e.Error(),
	)
}

func TestAMalformedIdentifierIsNotFoundDespiteUnrelatedItems(t *testing.T) {
	_, e := items(
		t,
		`{"Items":[{"Id":"0f0e0d0c0b0a09080706050403020100","Name":"delta"}],"TotalRecordCount":1}`,
	).Item("alfa")
	assert.True(t, not_found.Is(e))
}
