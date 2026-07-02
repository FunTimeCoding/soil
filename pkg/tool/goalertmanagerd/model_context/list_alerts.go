package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert_filter"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) listAlerts(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListAlerts,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	var f *alert_filter.Filter

	if a.Filter != "" || a.Active != "" || a.Silenced != "" ||
		a.Inhibited != "" || a.Unprocessed != "" || a.Receiver != "" {
		f = alert_filter.New()
		f.Receiver = a.Receiver

		if a.Filter != "" {
			f.Filter = strings.Split(a.Filter, ",")
		}

		if a.Active != "" {
			v := a.Active == "true"
			f.Active = &v
		}

		if a.Silenced != "" {
			v := a.Silenced == "true"
			f.Silenced = &v
		}

		if a.Inhibited != "" {
			v := a.Inhibited == "true"
			f.Inhibited = &v
		}

		if a.Unprocessed != "" {
			v := a.Unprocessed == "true"
			f.Unprocessed = &v
		}
	}

	v, stat, e := s.service.Alerts(instance, f)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		convert.AlertResult(v, stat, int(a.Limit), int(a.Offset)),
	)
}
