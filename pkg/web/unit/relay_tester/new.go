package relay_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net"
	"testing"
)

func New(
	t *testing.T,
	target string,
) *Relay {
	t.Helper()
	l, e := net.Listen(
		systemConstant.Transmission,
		net.JoinHostPort(constant.Loopback, "0"),
	)
	errors.PanicOnError(e)
	result := &Relay{listener: l, target: target}
	go result.accept()
	t.Cleanup(result.Close)

	return result
}
