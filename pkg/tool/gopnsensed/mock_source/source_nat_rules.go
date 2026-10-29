package mock_source

import "github.com/funtimecoding/soil/pkg/opnsense/source_nat"

func (c *Client) SourceNatRules(_ string) ([]*source_nat.Rule, error) {
	return nil, nil
}
