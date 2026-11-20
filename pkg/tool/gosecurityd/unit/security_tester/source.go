package security_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/tracker_source"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/mock_client"
	"path/filepath"
)

func (o *Tester) Source(
	c *mock_client.Client,
	cache string,
	supported bool,
	releases ...string,
) *tracker_source.Source {
	o.t.Helper()
	g := peer.New(
		cache,
		"tracker.test",
		c,
		logger.New(o.t.Context()),
		o.Reporter,
	)
	o.t.Cleanup(g.Close)

	return tracker_source.New(
		c,
		g,
		o.now,
		filepath.Join(o.Cache, cache),
		supported,
		releases...,
	)
}
