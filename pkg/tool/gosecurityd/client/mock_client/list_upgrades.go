package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (c *Client) ListUpgrades(
	_ context.Context,
	p *client.ListUpgradesParams,
) (*client.UpgradeList, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	show := ""
	lowest := 0.0

	if p != nil && p.Show != nil {
		show = *p.Show
	}

	if p != nil {
		lowest = convert.Minimum(p.Score)
	}

	result := &client.UpgradeList{
		Groups:   []client.UpgradeGroup{},
		Releases: []string{},
	}
	hosts := map[string]bool{}

	for _, g := range c.upgrades {
		if g.Kernel != (show == constant.KernelFilter) ||
			(show != "" && show != constant.KernelFilter && g.Release != show) ||
			g.TopScore < lowest {
			continue
		}

		result.Groups = append(result.Groups, g)
		result.Vulnerabilities += len(g.Vulnerabilities)

		for _, h := range g.Hosts {
			hosts[h.Name] = true
		}
	}

	result.Hosts = len(hosts)

	return result, nil
}
