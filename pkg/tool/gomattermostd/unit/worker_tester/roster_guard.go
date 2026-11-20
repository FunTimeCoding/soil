package worker_tester

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
)

func (o *Tester) RosterGuard(
	c goclauded.Client,
	r soilFace.Reporter,
) *peer.Guard {
	o.t.Helper()
	result := peer.New(
		constant.RosterPeer,
		c.Address(),
		c,
		logger.New(o.t.Context()),
		r,
	)
	o.t.Cleanup(result.Close)

	return result
}
