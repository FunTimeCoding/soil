package rejecting_host

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
)

func (c *Client) AddHost(_ *request.Host) (string, error) {
	return "", validation.New("host rejected: ip: invalid address")
}
