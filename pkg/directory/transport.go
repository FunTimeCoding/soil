package directory

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/system"
)

func (c *Client) transport() (*tls.Config, error) {
	result := &tls.Config{InsecureSkipVerify: c.untrusted}

	if c.authority == "" {
		return result, nil
	}

	pool := x509.NewCertPool()

	if !pool.AppendCertsFromPEM(system.ReadBytesUnsafe(c.authority)) {
		return nil, not_configured.Format(
			"%s authority holds no certificate: %s",
			constant.Subject,
			c.authority,
		)
	}

	result.RootCAs = pool

	return result, nil
}
