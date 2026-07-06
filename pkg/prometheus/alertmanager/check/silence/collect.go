package silence

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
)

func collect(c *alertmanager.Client) []*silence.Silence {
	return c.MustSilences(true)
}
