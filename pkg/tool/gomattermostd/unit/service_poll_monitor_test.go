package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/counting_mattermost"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"testing"
)

func TestFailingFirstPollAnswersItsFault(t *testing.T) {
	o := service_tester.New(t, service_tester.WithMonitor())
	s := o.Service()
	o.Mattermost.Failure = service_tester.Fault()
	failures := s.PollMonitor()
	assert.Count(t, 1, failures)
	assert.True(t, classify.Reportable(failures[0]))
}

func TestMattermostOutageAtPollIsLeftToGuard(t *testing.T) {
	o := service_tester.New(t, service_tester.WithMonitor())
	s := o.Service()
	o.Mattermost.Failure = unavailable.Format("mattermost maintenance")
	failures := s.PollMonitor()
	assert.Count(t, 1, failures)
	assert.False(t, classify.Reportable(failures[0]))
}

func TestPollInitializesOnceMattermostAnswers(t *testing.T) {
	m := service_tester.New(t).Mattermost
	c := counting_mattermost.New(m)
	defer c.Close()
	s := service_tester.New(
		t,
		service_tester.WithMonitor(),
		service_tester.WithMattermost(c),
	).Service()
	m.Failure = service_tester.Fault()
	assert.Count(t, 1, s.PollMonitor())
	m.Failure = nil
	assert.Count(t, 0, s.PollMonitor())
	assert.Count(t, 0, s.PollMonitor())
	assert.Integer(t, 2, c.Mes())
	assert.Integer(t, 1, c.TeamChannels())
}
