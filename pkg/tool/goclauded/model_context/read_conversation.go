package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) readConversation(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadConversation,
) (*mcp.CallToolResult, error) {
	if _, e := s.resolveCaller(x, constant.ReadConversation); e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	blocks, e := s.service.ReadConversation(a.Session, a.Around, int(a.Count))

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if len(blocks) == 0 {
		return response.Fail(
			"no block %s in conversation %s",
			a.Around,
			a.Session,
		)
	}

	var lines []string

	for _, b := range blocks {
		header := fmt.Sprintf(
			"[%s %s %s %s]",
			b.At,
			b.Role,
			b.Kind,
			b.Identifier,
		)

		if b.Identifier == a.Around {
			header = join.Space(header, "<- hit")
		}

		lines = append(lines, header, b.Text, "")
	}

	return response.Success(join.NewLine(lines))
}
