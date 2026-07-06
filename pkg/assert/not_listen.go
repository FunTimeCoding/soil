package assert

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
	"testing"
)

func NotListen(
	t *testing.T,
	port int,
) {
	t.Helper()
	c, e := net.DialTimeout(
		constant.Transmission,
		fmt.Sprintf("localhost:%d", port),
		constant.Retry,
	)

	if e == nil {
		if f := c.Close(); f != nil {
			panic(f)
		}

		t.Errorf("port %d is listening", port)
	}
}
