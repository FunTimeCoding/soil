package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/failing_jira"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	"testing"
)

func TestTheServiceClosesBothClientsItBuilt(t *testing.T) {
	j := service_tester.NewJira(t)
	c := service_tester.NewConfluence(t)
	service.New(
		service_tester.NewOption(j, c, nil),
		logger.New(t.Context()),
		memory.New(),
	).Close()
	assert.Integer(t, 1, j.Closed)
	assert.Integer(t, 1, c.Closed)
}

func TestAPollFillsThePlateFromTheService(t *testing.T) {
	j := service_tester.NewJira(t)
	j.Issues = []*issue.Issue{
		issue.New(issue.Raw("ALFA-1"), issue.FixtureOption()),
	}
	k := worker.New(
		service_tester.NewService(t, j, service_tester.NewConfluence(t), nil),
		constant.PollInterval,
		logger.New(context.Background()),
		memory.New(),
	)
	defer k.Close()
	k.Poll()
	assert.Count(t, 1, k.Issues())
	assert.String(t, "ALFA-1", k.Issues()[0].Key)
}

func TestAFailedPollIsReportedAndKeepsThePlate(t *testing.T) {
	r := memory.New()
	k := worker.New(
		service_tester.NewService(
			t,
			failing_jira.New(errors.Format("search refused", "ALFA")),
			service_tester.NewConfluence(t),
			nil,
		),
		constant.PollInterval,
		logger.New(context.Background()),
		r,
	)
	defer k.Close()
	k.Poll()
	assert.Count(t, 1, r.Events())
	assert.Count(t, 0, k.Issues())
}

func TestAnOutageLeavesTheReportToTheGuard(t *testing.T) {
	r := memory.New()
	k := worker.New(
		service_tester.NewService(
			t,
			failing_jira.New(refusedConnection()),
			service_tester.NewConfluence(t),
			nil,
		),
		constant.PollInterval,
		logger.New(context.Background()),
		r,
	)
	defer k.Close()
	k.Poll()
	assert.Count(t, 0, r.Events())
}
