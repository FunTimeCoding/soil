package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
)

func (o *Tester) Guard(
	c face.MattermostClient,
	r soilFace.Reporter,
) *peer.Guard {
	o.t.Helper()
	result := peer.New(
		constant.MattermostPeer,
		c.Host(),
		c,
		logger.New(o.t.Context()),
		r,
	)
	o.t.Cleanup(result.Close)

	return result
}
