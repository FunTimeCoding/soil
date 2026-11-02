package mock_client

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DeriveHardwareAddress(
	_ context.Context,
	p *client.DeriveHardwareAddressParams,
) (*client.DerivedAddress, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := &client.DerivedAddress{
		Identifier:      p.Identifier,
		HardwareAddress: fmt.Sprintf(
			"02:00:00:00:%02x:%02x",
			p.Identifier/256,
			p.Identifier%256,
		),
	}

	if p.Instance != nil {
		result.Instance = *p.Instance
	}

	return result, nil
}
