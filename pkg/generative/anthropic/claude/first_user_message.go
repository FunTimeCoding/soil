package claude

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func (c *Client) FirstUserMessage(sessionIdentifier string) string {
	path := filepath.Join(
		c.base,
		join.Empty(sessionIdentifier, constant.NotationLogExtension),
	)
	f, e := os.Open(path)

	if e != nil {
		return ""
	}

	defer errors.PanicClose(f)
	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 1024*1024), 1024*1024)

	for s.Scan() {
		var l notation.Line

		if json.Unmarshal(s.Bytes(), &l) != nil {
			continue
		}

		if l.Type != "user" || l.Meta {
			continue
		}

		if l.Message == nil {
			continue
		}

		var m notation.Message

		if json.Unmarshal(l.Message, &m) != nil {
			continue
		}

		text := ExtractText(m.Content)

		if IsSystemNoise(text) {
			continue
		}

		clean := CleanContent(text)

		if len(clean) > 30 {
			return clean
		}
	}

	return ""
}
