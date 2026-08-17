package tracker

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"io"
	"os"
)

func Read(
	path string,
	s *State,
	followed []string,
) ([]tool_call.Call, error) {
	f, e := os.Open(path)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(f)
	i, e := f.Stat()

	if e != nil {
		return nil, e
	}

	if s.Offset > i.Size() {
		s.Reset()
	}

	cold := s.Offset == 0

	if !cold {
		_, e = f.Seek(s.Offset, io.SeekStart)

		if e != nil {
			return nil, e
		}
	}

	var calls []tool_call.Call
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

			if json.Unmarshal(line.Message, &m) != nil {
				continue
			}

			if m.Usage != nil {
				s.recordUsage(line.Request, &m)
			}

			for _, b := range blocks(m.Content) {
				if b.Type == constant.ClaudeToolUseBlock &&
					follow(b.Name, followed) {
					s.recordCall(&b, line.Timestamp)
				}
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

		for _, b := range blocks(m.Content) {
			if b.Type != constant.ClaudeToolResultBlock {
				continue
			}

			if c := s.resolveCall(&b); c != nil {
				calls = append(calls, *c)
			}
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
		return nil, e
	}

	s.Offset = offset

	return calls, scanner.Err()
}
