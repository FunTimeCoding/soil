package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
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
		return response.Success("No usage data yet.")
	}

	lines := []string{
		fmt.Sprintf(
			"Session  %2d%%   resets %s",
			result.FiveHourPercent,
			result.FiveHourResetText(),
		),
		fmt.Sprintf(
			"Weekly   %2d%%   resets %s",
			result.SevenDayPercent,
			result.SevenDayResetText(),
		),
	}

	if result.HasFable() {
		lines = append(
			lines,
			fmt.Sprintf(
				"Fable    %2d%%   resets %s",
				result.FablePercent,
				result.FableReset,
			),
		)
	}

	lines = append(
		lines,
		fmt.Sprintf("Updated  %s", time.FormatCompact(result.LastUpdated)),
	)

	return response.Success(join.NewLine(lines))
}
