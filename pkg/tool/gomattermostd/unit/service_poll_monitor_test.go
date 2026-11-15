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
	m := service_tester.NewMattermost(t, "kilo")
	s := service_tester.Monitor(t, m)
	m.Failure = service_tester.Fault()
	failures := s.PollMonitor()
	assert.Count(t, 1, failures)
	assert.True(t, classify.Reportable(failures[0]))
}

func TestMattermostOutageAtPollIsLeftToGuard(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	s := service_tester.Monitor(t, m)
	m.Failure = unavailable.Format("mattermost maintenance")
	failures := s.PollMonitor()
	assert.Count(t, 1, failures)
	assert.False(t, classify.Reportable(failures[0]))
}

func TestPollInitializesOnceMattermostAnswers(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	c := counting_mattermost.New(m)
	defer c.Close()
	s := service_tester.Monitor(t, c)
	m.Failure = service_tester.Fault()
	assert.Count(t, 1, s.PollMonitor())
	m.Failure = nil
	assert.Count(t, 0, s.PollMonitor())
	assert.Count(t, 0, s.PollMonitor())
	assert.Integer(t, 2, c.Mes())
	assert.Integer(t, 1, c.TeamChannels())
}
