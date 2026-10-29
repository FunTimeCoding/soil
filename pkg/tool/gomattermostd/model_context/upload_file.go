package model_context

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"io/fs"
	"os"
	"path/filepath"
)

func (s *Server) UploadFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UploadFile,
) (*mcp.CallToolResult, error) {
	if a.Path == "" {
		return response.Fail("path is required")
	}

	if a.ChannelIdentifier == "" && a.ChannelName == "" {
		return response.Fail("channel_id or channel_name is required")
	}

	b, e := os.ReadFile(a.Path)

	if errors.Is(e, fs.ErrNotExist) {
		return response.Fail("no such file: %s", a.Path)
	}

	if e != nil {
		return s.answerWith(e, "read file")
	}

	ch, e := s.service.ResolveChannel(a.ChannelIdentifier, a.ChannelName)

	if e != nil {
		return s.answerWith(e, "channel")
	}

	name := filepath.Base(a.Path)
	message := a.Message

	if message == "" {
		message = name
	}

	p, f, e := s.service.UploadAndPost(ch, name, b, message)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(
		map[string]any{
			"post_id":   p.Id,
			"file_id":   f.Id,
			"file_name": name,
			"channel":   s.service.ChannelDisplayName(ch),
			"message":   message,
			"create_at": formatMilli(p.CreateAt),
		},
	)
}
