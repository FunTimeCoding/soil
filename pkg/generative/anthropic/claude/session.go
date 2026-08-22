package claude

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
)

func (c *Client) Session(identifier string) *session.Session {
	return scanSession(
		filepath.Join(
			c.base,
			join.Empty(identifier, constant.NotationLogExtension),
		),
	)
}
