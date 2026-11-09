package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/subscribed"
	"testing"
)

func TestDigestWhoseNotifyFailedArrivesOnNextFlush(t *testing.T) {
	o := subscribed.New(t)
	o.Subscribe(t, "alfa", "kilo", "")
	o.Roster.Fail(not_found.New(constant.ParameterCallsign, "kilo"))
	o.Reply("bravo", "alfa", "first message")
	assert.Count(t, 1, o.Service.FlushDue())
	o.Roster.Fail(nil)
	assert.Count(t, 0, o.Service.FlushDue())
	assert.Count(t, 1, o.Roster.Notes())
	assert.StringContains(t, "first message", o.Roster.Notes()[0].Body)
	assert.Count(t, 0, o.Service.FlushDue())
	assert.Count(t, 1, o.Roster.Notes())
}

func TestPendingDigestCombinesWithNewerReplies(t *testing.T) {
	o := subscribed.New(t)
	o.Subscribe(t, "alfa", "kilo", "")
	o.Roster.Fail(not_found.New(constant.ParameterCallsign, "kilo"))
	o.Reply("bravo", "alfa", "first message")
	o.Service.FlushDue()
	o.Roster.Fail(nil)
	o.Reply("charlie", "alfa", "second message")
	assert.Count(t, 0, o.Service.FlushDue())
	assert.Count(t, 1, o.Roster.Notes())
	assert.StringContains(t, "2 new", o.Roster.Notes()[0].Body)
	assert.StringContains(t, "first message", o.Roster.Notes()[0].Body)
	assert.StringContains(t, "second message", o.Roster.Notes()[0].Body)
}

func TestFailingSubscriberIsReportedOncePerEpisode(t *testing.T) {
	o := subscribed.New(t)
	o.Subscribe(t, "alfa", "kilo", "")
	o.Roster.Fail(not_found.New(constant.ParameterCallsign, "kilo"))
	o.Reply("bravo", "alfa", "first message")
	assert.Count(t, 1, o.Service.FlushDue())
	assert.Count(t, 0, o.Service.FlushDue())
	assert.Count(t, 0, o.Service.FlushDue())
}

func TestOnlyFailedSubscriberIsRetried(t *testing.T) {
	o := subscribed.New(t)
	o.Subscribe(t, "alfa", "kilo", "")
	o.Subscribe(t, "alfa", "lima", "")
	o.Roster.FailFor("kilo", not_found.New(constant.ParameterCallsign, "kilo"))
	o.Reply("bravo", "alfa", "first message")
	o.Service.FlushDue()
	o.Roster.FailFor("kilo", nil)
	o.Service.FlushDue()
	var kilo, lima int

	for _, v := range o.Roster.Notes() {
		switch v.Callsign {
		case "kilo":
			kilo++
		case "lima":
			lima++
		}
	}

	assert.Integer(t, 1, kilo)
	assert.Integer(t, 1, lima)
}

func TestPendingDigestOfUnsubscribedThreadIsDropped(t *testing.T) {
	o := subscribed.New(t)
	o.Subscribe(t, "alfa", "kilo", "")
	o.Roster.Fail(not_found.New(constant.ParameterCallsign, "kilo"))
	o.Reply("bravo", "alfa", "first message")
	o.Service.FlushDue()
	_, e := o.Service.Unsubscribe("kilo", "alfa")
	assert.FatalOnError(t, e)
	o.Roster.Fail(nil)
	assert.Count(t, 0, o.Service.FlushDue())
	assert.Count(t, 0, o.Roster.Notes())
}
