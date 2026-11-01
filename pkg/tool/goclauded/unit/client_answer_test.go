package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
	"time"
)

func TestTheClientAnswersNoUsageAsAbsent(t *testing.T) {
	s := base.New(t)
	r, found, e := s.Connector(t).GetUsage(t.Context())
	assert.FatalOnError(t, e)
	assert.False(t, found)
	assert.True(t, r == nil)
}

func TestTheClientReadsRecordedUsage(t *testing.T) {
	s := base.New(t)
	now := s.Store.Clock()()
	assert.FatalOnError(
		t,
		s.Service.RecordRateLimits(0, 20, time.Time{}, now.Add(24*time.Hour)),
	)
	r, found, e := s.Connector(t).GetUsage(t.Context())
	assert.FatalOnError(t, e)
	assert.True(t, found)
	assert.Integer(t, 20, r.SevenDayPercent)
}

func TestTheClientAnswersAReleasedCallsignHoldAsAbsent(t *testing.T) {
	s := base.NewWithHold(t, constant.FixtureChannelHold)
	s.Check("session-1")
	s.Store.Advance(time.Minute)
	r, found, e := s.Connector(t).GetChannelCallsign(
		t.Context(),
		&client.GetChannelCallsignParams{
			Session: "session-1",
			Since:   s.Store.Clock()(),
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, found)
	assert.True(t, r == nil)
}

func TestTheClientReturnsTheMatchesOfAnAmbiguousQuery(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("abc-1111")
	s.Store.EnsureSession("abc-2222")
	r, matches, e := s.Connector(t).GetResolve(
		t.Context(),
		&client.GetResolveParams{Query: "abc"},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r == nil)
	assert.Count(t, 2, matches)
}

func TestTheClientResolvesAUniqueQuery(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("abc-1111")
	r, matches, e := s.Connector(t).GetResolve(
		t.Context(),
		&client.GetResolveParams{Query: "abc"},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, matches)
	assert.String(t, "abc-1111", r.Identifier)
}
