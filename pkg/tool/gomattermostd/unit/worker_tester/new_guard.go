package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"testing"
)

func NewGuard(
	t *testing.T,
	c face.Client,
	r soilFace.Reporter,
) *peer.Guard {
	t.Helper()
	result := peer.New(
		constant.MattermostPeer,
		c.Host(),
		c,
		logger.New(t.Context()),
		r,
	)
	t.Cleanup(result.Close)

	return result
}
