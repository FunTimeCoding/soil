package worker_tester

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"testing"
)

func NewRosterGuard(
	t *testing.T,
	c goclauded.Client,
	r soilFace.Reporter,
) *peer.Guard {
	t.Helper()
	result := peer.New(
		constant.RosterPeer,
		c.Address(),
		c,
		logger.New(t.Context()),
		r,
	)
	t.Cleanup(result.Close)

	return result
}
