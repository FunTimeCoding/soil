package assert

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net"
	"strconv"
	"sync"
	"testing"
)

func SkipUnlessDialable(
	t testing.TB,
	host string,
) {
	t.Helper()
	address := host

	if _, _, e := net.SplitHostPort(host); e != nil {
		address = net.JoinHostPort(host, strconv.Itoa(webConstant.SecurePort))
	}

	probe, _ := dialed.LoadOrStore(
		host,
		sync.OnceValue(
			func() error {
				c, e := net.DialTimeout(
					systemConstant.Transmission,
					address,
					webConstant.ProbeDialTimeout,
				)

				if e != nil {
					return e
				}

				return c.Close()
			},
		),
	)

	if e := probe.(func() error)(); e != nil {
		t.Skip(e)
	}
}
