package tracker

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"io"
	"os"
)

func Read(
	path string,
	s *State,
) error {
	f, e := os.Open(path)

	if e != nil {
		return e
	}

	defer errors.PanicClose(f)
	i, e := f.Stat()

	if e != nil {
		return e
	}

	if s.Offset > i.Size() {
		s.Reset()
	}

	cold := s.Offset == 0

	if !cold {
		_, e = f.Seek(s.Offset, io.SeekStart)

		if e != nil {
			return e
		}
	}

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	for scanner.Scan() {
		s.Lines++
		var line notation.Line

		if json.Unmarshal(scanner.Bytes(), &line) != nil {
			continue
		}

		if line.Timestamp != "" {
			s.LastTimestamp = line.Timestamp

			if s.FirstTimestamp == "" {
				s.FirstTimestamp = line.Timestamp
			}
		}

		if cold {
			if line.Slug != "" && s.Slug == "" {
				s.Slug = line.Slug
			}

			if line.WorkDirectory != "" && s.WorkDirectory == "" {
				s.WorkDirectory = line.WorkDirectory
			}

			if line.Branch != "" && s.Branch == "" {
				s.Branch = line.Branch
			}
		}

		if line.Type == "assistant" && line.Message != nil {
			var m notation.Message

			if json.Unmarshal(line.Message, &m) == nil &&
				m.Usage != nil {
				s.recordUsage(line.Request, &m)
			}

			continue
		}

		if line.Type != "user" || line.Meta || line.Message == nil {
			continue
		}

		var m notation.Message

		if json.Unmarshal(line.Message, &m) != nil {
			continue
		}

		text := claude.ExtractText(m.Content)

		if claude.IsSystemNoise(text) {
			continue
		}

		clean := claude.CleanContent(text)

		if len(clean) <= 20 {
			continue
		}

		s.UserMessageCount++

		if s.FirstMessage == "" {
			if len(clean) > 80 {
				clean = clean[:80]
			}

			s.FirstMessage = clean
		}
	}

	offset, e := f.Seek(0, io.SeekCurrent)

	if e != nil {
		return e
	}

	s.Offset = offset

	return scanner.Err()
}
