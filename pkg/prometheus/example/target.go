package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/prometheus/common/model"
)

func Target() {
	c := prometheus.NewEnvironment()
	r := c.MustTargets()

	for _, t := range r.Active {
		console.Format("Active: %+v\n", t.ScrapePool)
	}

	for _, t := range r.Dropped {
		address := t.DiscoveredLabels[model.AddressLabel]
		console.Format("Dropped: %+v\n", address)
	}
}
