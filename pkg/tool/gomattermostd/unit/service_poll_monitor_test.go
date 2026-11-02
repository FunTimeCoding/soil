package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"testing"
)

func TestAFailingFirstPollAnswersItsFault(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	s := monitorService(t, m)
	m.Failure = fault()
	failures := s.PollMonitor()
	assert.Count(t, 1, failures)
	assert.True(t, classify.Reportable(failures[0]))
}

func TestAMattermostOutageAtPollIsLeftToTheGuard(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	s := monitorService(t, m)
	m.Failure = unavailable.Format("mattermost maintenance")
	failures := s.PollMonitor()
	assert.Count(t, 1, failures)
	assert.False(t, classify.Reportable(failures[0]))
}

func TestThePollInitializesOnceMattermostAnswers(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	s := monitorService(t, m)
	m.Failure = fault()
	assert.Count(t, 1, s.PollMonitor())
	m.Failure = nil
	assert.Count(t, 0, s.PollMonitor())
}
