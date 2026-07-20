package service

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"os"
	"strings"
	"time"
)

func (s *Service) UsageBetween(
	from time.Time,
	to time.Time,
) map[string]*pricing.Tokens {
	entries, e := os.ReadDir(s.harbor)

	if e != nil {
		s.reporter.CaptureException(e)

		return nil
	}

	harbor := claude.NewDirectory(s.harbor)
	result := map[string]*pricing.Tokens{}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(
			entry.Name(),
			constant.NotationLogExtension,
		) {
			continue
		}

		i, f := entry.Info()

		if f != nil || i.ModTime().Before(from) {
			continue
		}

		identifier := strings.TrimSuffix(
			entry.Name(),
			constant.NotationLogExtension,
		)

		for _, u := range harbor.UsageEntries(identifier) {
			t, g := time.Parse(time.RFC3339Nano, u.Timestamp)

			if g != nil || t.Before(from) || !t.Before(to) {
				continue
			}

			key := pricing.NormalizeModel(u.Model)
			bucket, okay := result[key]

			if !okay {
				bucket = pricing.New()
				result[key] = bucket
			}

			bucket.AddEntry(u)
		}
	}

	return result
}
