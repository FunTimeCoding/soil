package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/prometheus/rule"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) queryRules(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.QueryRules,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	v, e := s.service.Rules(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	ruleType := a.Type

	if ruleType == "" {
		ruleType = rule.AlertType
	}

	switch ruleType {
	case rule.AlertType, rule.RecordType, rule.AllType:
	default:
		return response.Fail(
			"invalid type: %s (expected alert, record, or all)",
			ruleType,
		)
	}

	var filtered []*rule.Rule

	for _, r := range v {
		if ruleType == rule.AlertType && !r.IsAlert() {
			continue
		}

		if ruleType == rule.RecordType && !r.IsRecord() {
			continue
		}

		if a.Name != "" && !strings.Contains(r.Name, a.Name) {
			continue
		}

		if a.Group != "" && !strings.Contains(r.Group, a.Group) {
			continue
		}

		if a.State != "" && r.State != a.State {
			continue
		}

		filtered = append(filtered, r)
	}

	return response.SuccessAny(
		convert.NewRuleResult(filtered, int(a.Limit), int(a.Offset)),
	)
}
