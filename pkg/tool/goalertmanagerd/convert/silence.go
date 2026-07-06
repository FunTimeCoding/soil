package convert

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"time"
)

func Silence(s *silence.Silence) *SlimSilence {
	var start, end string

	if s.Start != nil {
		start = s.Start.Format(time.RFC3339)
	}

	if s.End != nil {
		end = s.End.Format(time.RFC3339)
	}

	var remaining string

	if r := s.Remain(); r > 0 {
		remaining = r.Round(time.Second).String()
	}

	return &SlimSilence{
		Identifier: s.Identifier,
		State:      s.State,
		Rule:       s.Rule,
		Match:      s.Match,
		Author:     s.Author,
		Comment:    s.Comment,
		Start:      start,
		End:        end,
		Remaining:  remaining,
		Link:       s.Link,
	}
}
