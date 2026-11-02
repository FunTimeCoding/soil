package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/unit/channel_tester"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestChannelConfirmRejectsWrongNonce(t *testing.T) {
	s := channel_tester.NewChannelServer()
	assert.True(
		t,
		channel_tester.Confirm(s, "not-the-nonce", stringsConstant.UpperAlfa).IsError,
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
	s.Resolve(stringsConstant.UpperAlfa)
	assert.True(
		t,
		channel_tester.Confirm(s, s.Nonce(), stringsConstant.UpperBravo).IsError,
	)
	assert.False(t, s.Opened())
}

func TestChannelConfirmOpensWhenCallsignMatchesResolved(t *testing.T) {
	s := channel_tester.NewChannelServer()
	s.Resolve(stringsConstant.UpperAlfa)
	assert.False(
		t,
		channel_tester.Confirm(s, s.Nonce(), stringsConstant.UpperAlfa).IsError,
	)
	assert.True(t, s.Opened())
}

func TestChannelConfirmAdoptsCallsignWhenUnresolved(t *testing.T) {
	s := channel_tester.NewChannelServer()
	assert.False(
		t,
		channel_tester.Confirm(s, s.Nonce(), stringsConstant.UpperAlfa).IsError,
	)
	assert.True(t, s.Opened())
	assert.True(
		t,
		channel_tester.Confirm(s, s.Nonce(), stringsConstant.UpperBravo).IsError,
	)
}
