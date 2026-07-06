package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/time"
	"os"
	"path/filepath"
	"strings"
)

func (s *Server) exportSession(
	i *session.Session,
	basePath string,
) string {
	o, f := s.service.GetSession(i.Identifier)
	errors.PanicOnError(f)
	alias := ""

	if o != nil {
		alias = o.AliasValue()
	}

	messages := s.service.Messages(i.Identifier)
	created := parseTimestamp(i.Timestamp)
	d := filepath.Join(
		basePath,
		created.Format("2006"),
		created.Format("01"),
	)
	slug := alias

	if slug == "" {
		slug = i.Slug
	}

	if slug == "" {
		slug = i.Identifier[:8]
	}

	name := fmt.Sprintf(
		"%s-%s.md",
		created.Format(time.DateYear),
		strings.ReplaceAll(
			strings.ToLower(slug),
			" ",
			"-",
		),
	)
	var b strings.Builder
	writer.Print(&b, "---\n")
	writer.Print(&b, "session: %s\n", i.Identifier)

	if alias != "" {
		writer.Print(&b, "alias: %s\n", alias)
	}

	if i.Slug != "" {
		writer.Print(&b, "slug: %s\n", i.Slug)
	}

	writer.Print(
		&b,
		"created: %s\n",
		created.Format(time.DateYear),
	)
	writer.Print(&b, "---\n\n")

	for _, m := range messages {
		if m.Text == "" || m.IsMeta {
			continue
		}

		writer.Print(&b, "## %s\n\n%s\n\n", m.Role, m.Text)
	}

	system.MakeDirectory(d)
	path := filepath.Join(d, name)
	e := os.WriteFile(path, []byte(b.String()), 0644)

	if e != nil {
		panic(e)
	}

	return path
}
