package outage_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"testing"
)

func Guard(
	t *testing.T,
	workers []face.Worker,
	name string,
) *peer.Guard {
	t.Helper()

	for _, w := range workers {
		if g, okay := w.(*peer.Guard); okay && g != nil && g.Name() == name {
			return g
		}
	}

	t.Fatalf("no guard for peer %s", name)

	return nil
}
