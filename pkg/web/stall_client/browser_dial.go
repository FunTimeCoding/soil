package stall_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/refraction-networking/utls"
	"net"
)

func browserDial(
	d *net.Dialer,
	insecure bool,
) func(context.Context, string, string) (net.Conn, error) {
	return func(
		x context.Context,
		network string,
		address string,
	) (net.Conn, error) {
		host, _, e := net.SplitHostPort(address)

		if e != nil {
			return nil, e
		}

		raw, e := d.DialContext(x, network, address)

		if e != nil {
			return nil, e
		}

		c := tls.UClient(
			raw,
			&tls.Config{ServerName: host, InsecureSkipVerify: insecure},
			tls.HelloChrome_Auto,
		)

		if f := c.BuildHandshakeState(); f != nil {
			errors.LogClose(raw)

			return nil, f
		}

		for _, extension := range c.Extensions {
			if a, okay := extension.(*tls.ALPNExtension); okay {
				a.AlpnProtocols = []string{constant.BrowserProtocol}
			}
		}

		if f := c.HandshakeContext(x); f != nil {
			errors.LogClose(raw)

			return nil, f
		}

		return c, nil
	}
}
