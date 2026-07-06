package claude

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"strings"
)

func (c *Client) Resolve(query string) *session.Session {
	if len(query) == 36 {
		path := filepath.Join(
			c.base,
			join.Empty(query, constant.NotationLogExtension),
		)
		s := scanSession(path)

		if s.Identifier != "" {
			return s
		}
	}

	for _, s := range c.Sessions() {
		if s.Identifier == query {
			return s
		}

		if s.Slug == query {
			return s
		}

		if strings.HasPrefix(s.Identifier, query) {
			return s
		}
	}

	return session.Stub()
}
