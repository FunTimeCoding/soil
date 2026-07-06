package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) tokenUsage(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	_, e := s.resolveCaller(x, constant.TokenUsage)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	result := s.service.Usage()

	if result == nil {
		return response.Success(
			"Usage monitoring not enabled or no data yet.",
		)
	}

	lines := []string{
		fmt.Sprintf(
			"Session    %2d%%   resets %s",
			result.SessionPercent,
			result.SessionReset,
		),
		fmt.Sprintf(
			"Weekly     %2d%%   resets %s",
			result.WeeklyAllPercent,
			result.WeeklyAllReset,
		),
		fmt.Sprintf("  Sonnet   %2d%%", result.WeeklySonnetPercent),
		fmt.Sprintf("  Design   %2d%%", result.WeeklyDesignPercent),
		fmt.Sprintf("Routines   %s", result.RoutineRuns),
		fmt.Sprintf(
			"Credits    %s   resets %s",
			result.CreditSpent,
			result.CreditReset,
		),
	}

	return response.Success(join.NewLine(lines))
}
