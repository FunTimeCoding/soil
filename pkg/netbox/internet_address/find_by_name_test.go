package internet_address

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFindByName(t *testing.T) {
	addresses := []*Address{{Name: fixtureAddress}}
	// Happy path
	assert.Any(
		t,
		&Address{Name: "192.168.0.1/24"},
		FindByName(addresses, fixtureAddress),
	)
	// Not found
	var expected *Address
	assert.Any(t, expected, FindByName(addresses, "192.168.0.2/24"))
}
