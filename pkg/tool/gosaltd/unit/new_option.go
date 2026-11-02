package unit

import (
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"path/filepath"
	"testing"
)

func newOption(t *testing.T) *option.Salt {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "gosaltd.db")
	result.Master = func() *salt.Client {
		return salt.New("localhost", 1, "guard", "guard")
	}

	return result
}
