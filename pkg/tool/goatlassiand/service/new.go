package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
)

func New(
	o *option.Atlassian,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	j := o.Jira()
	c := o.Confluence()

	return &Service{
		jira:            j,
		confluence:      c,
		jiraGuard:       peer.New(constant.JiraPeer, "", j, l, r),
		confluenceGuard: peer.New(constant.ConfluencePeer, "", c, l, r),
		project:         o.Project,
	}
}
