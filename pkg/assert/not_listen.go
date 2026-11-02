package assert

import (
	"fmt"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"net"
	"testing"
)

func NotListen(
	t *testing.T,
	port int,
) {
	t.Helper()
	c, e := net.DialTimeout(
		systemConstant.Transmission,
		fmt.Sprintf("localhost:%d", port),
		systemConstant.Retry,
	)

	if e == nil {
		if f := c.Close(); f != nil {
			panic(f)
		}

		t.Errorf("port %d is listening", port)
	}
}
