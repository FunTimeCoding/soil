package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/firefox"
	"github.com/funtimecoding/soil/pkg/firefox/unit/fake_extension"
	"testing"
	"time"
)

func connected(
	t *testing.T,
	c *firefox.Client,
	address string,
) *fake_extension.Extension {
	t.Helper()
	x := fake_extension.Connect(t, address)
	deadline := time.Now().Add(5 * time.Second)

	for {
		_, e := c.Tabs()

		if !unavailable.Is(e) {
			return x
		}

		if time.Now().After(deadline) {
			t.Fatal("extension never connected")
		}

		time.Sleep(10 * time.Millisecond)
	}
}
