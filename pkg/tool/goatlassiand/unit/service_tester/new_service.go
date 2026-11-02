package service_tester

import (
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"testing"
)

func NewService(
	t *testing.T,
	j atlassian.JiraClient,
	c atlassian.ConfluenceClient,
	project []string,
	n face.EventNotifier,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	r := memory.New()
	jiraGuard := peer.New(constant.JiraPeer, "", j, l, r)
	t.Cleanup(jiraGuard.Close)
	confluenceGuard := peer.New(constant.ConfluencePeer, "", c, l, r)
	t.Cleanup(confluenceGuard.Close)

	return service.New(j, c, jiraGuard, confluenceGuard, project, n)
}
