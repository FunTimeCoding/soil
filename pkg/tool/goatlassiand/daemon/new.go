package daemon

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web"
)

func New(
	o *option.Atlassian,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	j := o.Jira()
	c := o.Confluence()
	n := notifier.New()
	result := &Daemon{
		jira:            j,
		confluence:      c,
		jiraGuard:       peer.New(constant.JiraPeer, "", j, l, r),
		confluenceGuard: peer.New(constant.ConfluencePeer, "", c, l, r),
		logger:          l,
		reporter:        r,
	}
	result.service = service.New(
		j,
		c,
		result.jiraGuard,
		result.confluenceGuard,
		o.Project,
		n,
	)
	result.web = web.New(result.service, n)

	return result
}
