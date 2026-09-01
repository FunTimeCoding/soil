package client

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/kubernetes/filter"
)

func (c *Client) selectClients(f *filter.Filter) []*Client {
	var result []*Client

	if len(c.clients) == 0 {
		result = append(result, c)
	} else {
		for _, l := range c.clients {
			if c.Verbose {
				console.Format("select client: %s\n", l.cluster)
				console.Format("filter: %+v\n", f)
			}

			if f == nil || f.ContainsCluster(l.cluster) {
				result = append(result, l)
			}
		}
	}

	return result
}
