package opnsense_tester

import (
	"github.com/funtimecoding/soil/pkg/opnsense"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *opnsense.Client {
	t.Helper()
	result := opnsense.New(web.TrimScheme(address), "key", "secret", true)
	t.Cleanup(result.Close)

	return result
}
