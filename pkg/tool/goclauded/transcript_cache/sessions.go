package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"sort"
	"strings"
)

func (c *Cache) Sessions() []*session.Session {
	files, e := filepath.Glob(
		filepath.Join(c.Base(), join.Empty("*", constant.NotationLogExtension)),
	)

	if e != nil {
		return nil
	}

	var result []*session.Session

	for _, f := range files {
		identifier := strings.TrimSuffix(
			filepath.Base(f),
			constant.NotationLogExtension,
		)
		s := c.cachedSession(identifier)

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
