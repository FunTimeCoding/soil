package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	l := logger.New(o.t.Context())
	jiraGuard := peer.New(constant.JiraPeer, "", o.Jira, l, o.Reporter)
	o.t.Cleanup(jiraGuard.Close)
	confluenceGuard := peer.New(
		constant.ConfluencePeer,
		"",
		o.Confluence,
		l,
		o.Reporter,
	)
	o.t.Cleanup(confluenceGuard.Close)
	o.service = service.New(
		o.Jira,
		o.Confluence,
		jiraGuard,
		confluenceGuard,
		o.project,
		o.notifier,
	)

	return o.service
}
