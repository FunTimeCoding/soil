package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/unit/channel_tester"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestChannelConfirmRejectsWrongNonce(t *testing.T) {
	s := channel_tester.NewChannelServer()
	assert.True(
		t,
		channel_tester.Confirm(s, "not-the-nonce", constant.UpperAlfa).IsError,
	)
	assert.False(t, s.Opened())
}

func TestChannelConfirmRequiresCallsign(t *testing.T) {
	s := channel_tester.NewChannelServer()
	assert.True(t, channel_tester.Confirm(s, s.Nonce(), "").IsError)
	assert.False(t, s.Opened())
}

func TestChannelConfirmRejectsCallsignMismatch(t *testing.T) {
	s := channel_tester.NewChannelServer()
	s.Resolve(constant.UpperAlfa)
	assert.True(
		t,
		channel_tester.Confirm(s, s.Nonce(), constant.UpperBravo).IsError,
	)
	assert.False(t, s.Opened())
}

func TestChannelConfirmOpensWhenCallsignMatchesResolved(t *testing.T) {
	s := channel_tester.NewChannelServer()
	s.Resolve(constant.UpperAlfa)
	assert.False(
		t,
		channel_tester.Confirm(s, s.Nonce(), constant.UpperAlfa).IsError,
	)
	assert.True(t, s.Opened())
}

func TestChannelConfirmAdoptsCallsignWhenUnresolved(t *testing.T) {
	s := channel_tester.NewChannelServer()
	assert.False(
		t,
		channel_tester.Confirm(s, s.Nonce(), constant.UpperAlfa).IsError,
	)
	assert.True(t, s.Opened())
	assert.True(
		t,
		channel_tester.Confirm(s, s.Nonce(), constant.UpperBravo).IsError,
	)
}
