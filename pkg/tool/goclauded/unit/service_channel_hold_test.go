package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
	"time"
)

func TestChannelHoldEndsBeforeClientGivesUp(t *testing.T) {
	assert.True(t, constant.ChannelHold < webConstant.LongResponseHeaderTimeout)
}

func TestAwaitImmediateQueueReleasesAfterHold(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	drained, e := s.Message.AwaitImmediateQueue(
		context.Background(),
		"session-1",
		r1.Callsign,
		20*time.Millisecond,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
}

func TestAwaitCallsignReleasesAfterHold(t *testing.T) {
	s := service_tester.New(t)
	callsign, e := s.Message.AwaitCallsign(
		context.Background(),
		"session-unknown",
		s.Store.Clock()(),
		20*time.Millisecond,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "", callsign)
}
