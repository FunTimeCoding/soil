package mock_client

import (
	"context"
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

	if p == nil {
		p = &client.ListUpgradesParams{}
	}

	s, e := convert.Selection(p.Show, p.Label, p.Search, nil, 0)

	if e != nil {
		return nil, e
	}

	release := ""

	if p.Release != nil {
		release = *p.Release
	}

	kernel := p.Kernel != nil && *p.Kernel
	lowest := convert.Minimum(p.Score)
	result := &client.UpgradeList{
		Groups:   []client.UpgradeGroup{},
		Releases: []string{},
	}
	hosts := map[string]bool{}

	for _, g := range c.upgrades {
		if g.Kernel != kernel ||
			(release != "" && g.Release != release) ||
			g.TopScore < lowest {
			continue
		}

		if narrows(s) {
			kept := []client.UpgradeHost{}

			for _, h := range g.Hosts {
				if c.chosen(h.Name, s) {
					kept = append(kept, h)
				}
			}

			if len(kept) == 0 {
				continue
			}

			g.Hosts = kept
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
