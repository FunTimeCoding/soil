package assert

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
	"testing"
	"time"
)

func Listen(
	t *testing.T,
	port int,
) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) {
		c, e := net.DialTimeout(
			constant.Transmission,
			fmt.Sprintf("localhost:%d", port),
			50*time.Millisecond,
		)

		if e == nil {
			if f := c.Close(); f != nil {
				panic(f)
			}

			return
		}

		time.Sleep(50 * time.Millisecond)
	}

	t.Errorf("port %d not listening", port)
}
