package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mattermost/mattermost/server/public/model"
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

	if a.ChannelID == "" && a.ChannelName == "" {
		return response.Fail(
			"channel_id or channel_name is required",
		)
	}

	b, e := os.ReadFile(a.Path)

	if e != nil {
		return s.captureFail(e, "file read failed")
	}

	var ch *model.Channel

	if a.ChannelName != "" {
		ch, e = s.client.TeamChannel(a.ChannelName)

		if e != nil {
			return s.captureFail(e, "channel not found")
		}
	} else {
		ch, e = s.client.Channel(a.ChannelID)

		if e != nil {
			return s.captureFail(e, "channel not found")
		}
	}

	upload, _, e := s.client.Nested().UploadFile(
		s.client.Context(),
		b,
		ch.Id,
		filepath.Base(a.Path),
	)

	if e != nil {
		return s.captureDetail(e)
	}

	if len(upload.FileInfos) == 0 {
		return response.Fail("upload returned no file info")
	}

	fileIdentifier := upload.FileInfos[0].Id
	message := a.Message

	if message == "" {
		message = filepath.Base(a.Path)
	}

	p, e := s.client.Post(
		&model.Post{
			ChannelId: ch.Id,
			Message:   message,
			FileIds:   []string{fileIdentifier},
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	channelName := s.channelDisplayName(ch)

	return response.SuccessAny(
		map[string]any{
			"post_id":   p.Id,
			"file_id":   fileIdentifier,
			"file_name": filepath.Base(a.Path),
			"channel":   channelName,
			"message":   message,
			"create_at": formatMilli(p.CreateAt),
		},
	)
}
