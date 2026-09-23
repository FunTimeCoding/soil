package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel/mock_sink"
	"github.com/funtimecoding/soil/pkg/generative/unit/channel_tester"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
	"time"
)

func TestChannelAttachKnocksOnceWhenTheSessionAnswers(t *testing.T) {
	k := mock_sink.New()
	var s *channel.Server
	s = channel_tester.NewChannelServer().WithSink(k).WithSleep(
		func(_ time.Duration) {
			channel_tester.Confirm(s, s.Nonce(), stringConstant.UpperAlfa)
		},
	)
	s.Connected()
	s.Resolve(stringConstant.UpperAlfa)
	assert.String(t, "Alfa", s.Attach())
	assert.Count(t, 1, k.Kind(constant.ChannelAttachKind))
}

func TestChannelAttachCarriesTheNonceItWillAccept(t *testing.T) {
	k := mock_sink.New()
	var s *channel.Server
	s = channel_tester.NewChannelServer().WithSink(k).WithSleep(
		func(_ time.Duration) {
			channel_tester.Confirm(s, s.Nonce(), stringConstant.UpperAlfa)
		},
	)
	s.Connected()
	s.Resolve(stringConstant.UpperAlfa)
	s.Attach()
	attach := k.Kind(constant.ChannelAttachKind)
	assert.Count(t, 1, attach)
	assert.String(t, s.Nonce(), attach[0].Meta[constant.ChannelNonceMeta])
}

func TestChannelAttachStillKnocksWhenTheCallsignIsUnresolved(t *testing.T) {
	k := mock_sink.New()
	var s *channel.Server
	s = channel_tester.NewChannelServer().WithSink(k).WithSleep(
		func(_ time.Duration) {
			channel_tester.Confirm(s, s.Nonce(), stringConstant.UpperAlfa)
		},
	)
	s.Connected()
	s.Resolve("")
	assert.String(t, "Alfa", s.Attach())
	assert.Count(t, 1, k.Kind(constant.ChannelAttachKind))
}

func TestChannelAttachWaitsAMinuteBeforeKnockingAgain(t *testing.T) {
	k := mock_sink.New()
	var waits []time.Duration
	var s *channel.Server
	s = channel_tester.NewChannelServer().WithSink(k).WithSleep(
		func(d time.Duration) {
			waits = append(waits, d)
			channel_tester.Confirm(s, s.Nonce(), stringConstant.UpperAlfa)
		},
	)
	s.Connected()
	s.Resolve(stringConstant.UpperAlfa)
	s.Attach()
	assert.Count(t, 1, waits)
	assert.Duration(t, time.Minute, waits[0])
}

func TestChannelAttachBackoffDoublesToTheMaximum(t *testing.T) {
	k := mock_sink.New()
	var waits []time.Duration
	var s *channel.Server
	s = channel_tester.NewChannelServer().WithSink(k).WithSleep(
		func(d time.Duration) {
			waits = append(waits, d)

			if len(waits) == 6 {
				channel_tester.Confirm(s, s.Nonce(), stringConstant.UpperAlfa)
			}
		},
	)
	s.Connected()
	s.Resolve(stringConstant.UpperAlfa)
	s.Attach()
	assert.Duration(t, time.Minute, waits[0])
	assert.Duration(t, 2*time.Minute, waits[1])
	assert.Duration(t, 4*time.Minute, waits[2])
	assert.Duration(t, 5*time.Minute, waits[3])
	assert.Duration(t, 5*time.Minute, waits[5])
	assert.Count(t, 6, k.Kind(constant.ChannelAttachKind))
}
