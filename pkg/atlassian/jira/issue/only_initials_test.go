package issue

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/capital"
	"testing"
)

func TestOnlyInitials(t *testing.T) {
	assert.Any(
		t,
		[]*Issue{{Initials: "ALFA"}},
		OnlyInitials(
			[]*Issue{
				{Initials: capital.Alfa},
				{Initials: capital.Bravo},
			},
			capital.Alfa,
		),
	)
}
