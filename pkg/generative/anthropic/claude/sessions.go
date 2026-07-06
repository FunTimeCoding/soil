package claude

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"path/filepath"
	"sort"
)

func (c *Client) Sessions() []*session.Session {
	files, e := filepath.Glob(filepath.Join(c.base, "*.jsonl"))

	if e != nil {
		return nil
	}

	var result []*session.Session

	for _, f := range files {
		s := scanSession(f)

		if s.Identifier != "" {
			result = append(result, s)
		}
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Timestamp > result[j].Timestamp
		},
	)

	return result
}
