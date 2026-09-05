package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/remote"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestRemote(t *testing.T) {
	assert.Any(
		t,
		&remote.Remote{Name: "Alfa", Locator: "Bravo", Provider: "Charlie"},
		remote.New(
			constant.UpperAlfa,
			constant.UpperBravo,
			constant.UpperCharlie,
		),
	)
}
