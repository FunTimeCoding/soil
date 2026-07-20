package tracker

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *State) recordUsage(
	request string,
	m *notation.Message,
) {
	if m.Identifier != "" {
		key := join.Empty(m.Identifier, request)

		for _, seen := range s.RecentMessages {
			if seen == key {
				return
			}
		}

		s.RecentMessages = append(s.RecentMessages, key)

		if len(s.RecentMessages) > claude.RecentMessageLimit {
			s.RecentMessages = s.RecentMessages[1:]
		}
	}

	model := pricing.NormalizeModel(m.Model)

	if s.Usage == nil {
		s.Usage = map[string]*pricing.Tokens{}
	}

	t, okay := s.Usage[model]

	if !okay {
		t = pricing.New()
		s.Usage[model] = t
	}

	t.Input += m.Usage.InputTokens
	t.Output += m.Usage.OutputTokens
	t.CacheCreation += m.Usage.CacheCreationInputTokens
	t.CacheRead += m.Usage.CacheReadInputTokens

	if m.Usage.CacheCreation != nil {
		t.Cache5Minute += m.Usage.CacheCreation.Ephemeral5MinuteInputTokens
		t.Cache1Hour += m.Usage.CacheCreation.Ephemeral1HourInputTokens
	}

	t.Count++
}
