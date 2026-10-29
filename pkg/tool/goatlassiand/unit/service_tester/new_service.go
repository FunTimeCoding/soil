package service_tester

import (
	confluence "github.com/funtimecoding/soil/pkg/atlassian/confluence/face"
	jira "github.com/funtimecoding/soil/pkg/atlassian/jira/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"testing"
)

func NewService(
	t *testing.T,
	j jira.Client,
	c confluence.Client,
	project []string,
) *service.Service {
	t.Helper()
	result := service.New(
		NewOption(j, c, project),
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(result.Close)

	return result
}
