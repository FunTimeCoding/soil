package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestSetReplacesWholeSet(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.SetLabel("session-1", constant.FixtureLabelKey, "on_call,lead")
	s.SetLabel(
		"session-1",
		constant.FixtureLabelKey,
		constant.FixtureCompatible,
	)
	assert.String(t, "lead", labelled(t, s, constant.FixtureLabelKey))
}

func TestSetReadsBackSortedWithoutRepeats(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.SetLabel("session-1", constant.FixtureLabelKey, " on_call, lead,on_call,")
	assert.String(t, "lead,on_call", labelled(t, s, constant.FixtureLabelKey))
}

func TestEmptySetClearsKey(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.SetLabel("session-1", constant.FixtureLabelKey, "on_call,lead")
	s.SetLabel("session-1", constant.FixtureLabelKey, "")
	assert.String(t, "", labelled(t, s, constant.FixtureLabelKey))
}

func TestAddKeepsOtherValues(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.SetLabel("session-1", constant.FixtureLabelKey, "on_call")
	s.AddLabel(
		"session-1",
		constant.FixtureLabelKey,
		constant.FixtureCompatible,
	)
	assert.String(t, "lead,on_call", labelled(t, s, constant.FixtureLabelKey))
}

func TestRemoveTakesOnlyItsValue(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.SetLabel("session-1", constant.FixtureLabelKey, "on_call,lead")
	s.RemoveLabel(
		"session-1",
		constant.FixtureLabelKey,
		constant.FixtureCompatible,
	)
	assert.String(t, "on_call", labelled(t, s, constant.FixtureLabelKey))
}

func TestEventCarriesBothWholeSets(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.SetLabel("session-1", constant.FixtureLabelKey, "on_call")
	s.AddLabel(
		"session-1",
		constant.FixtureLabelKey,
		constant.FixtureCompatible,
	)
	events := s.Store.Events(event_query.New().Kind(constant.Label).SetLimit(1))
	assert.Count(t, 1, events)
	assert.String(t, "on_call", events[0].Metadata[constant.Past])
	assert.String(t, "lead,on_call", events[0].Metadata[constant.Now])
}

func TestAddingHeldValueLogsNothing(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.SetLabel(
		"session-1",
		constant.FixtureLabelKey,
		constant.FixtureCompatible,
	)
	answer := s.AddLabel(
		"session-1",
		constant.FixtureLabelKey,
		constant.FixtureCompatible,
	)
	assert.String(t, "role lead (unchanged)", answer)
	assert.Count(
		t,
		1,
		s.Store.Events(event_query.New().Kind(constant.Label).SetLimit(10)),
	)
}

func TestSurfaceSeesOneEntryPerKeyInCommaForm(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	s.SetLabel("session-1", constant.FixtureLabelKey, "on_call,lead")
	s.SetLabel("session-1", constant.BroadcastLabel, constant.BroadcastOff)
	labels, e := s.Label.LabelsBySession("session-1")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, labels)
	assert.String(t, "broadcast", labels[0].Key)
	assert.String(t, "off", labels[0].Value)
	assert.String(t, "role", labels[1].Key)
	assert.String(t, "lead,on_call", labels[1].Value)
}

func TestQuietSessionWithSeveralValuesStaysQuiet(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-2")
	s.Check("session-1")
	s.SetLabel("session-1", constant.BroadcastLabel, "off,muted")
	s.Send(r1.Callsign, "", "echo")
	assert.Integer(t, 0, s.Channel("session-1"))
	assert.Count(t, 0, s.Check("session-1").Entries)
}
