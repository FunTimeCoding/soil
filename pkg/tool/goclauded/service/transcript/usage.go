package transcript

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"github.com/funtimecoding/soil/pkg/time/window"
	"os"
	"strings"
	"time"
)

func (s *Service) Usage(days int) map[string]*pricing.Tokens {
	w := window.Whole().Bound(s.clock(), time.Duration(days)*24*time.Hour)
	from, _ := w.Start()
	to, _ := w.End()
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
			soil.NotationLogExtension,
		) {
			continue
		}

		i, f := entry.Info()

		if f != nil || i.ModTime().Before(from) {
			continue
		}

		identifier := strings.TrimSuffix(
			entry.Name(),
			soil.NotationLogExtension,
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
