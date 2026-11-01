package session_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/card"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func rosterCard(
	s *session.Session,
	labels []label.Label,
	latest *pulse.Pulse,
) gomponents.Node {
	result := card.New(s.Name).
		Link(Path(s.Identifier)).
		Marker(component.ActivityDot(activity(s)))

	if s.Topic != "" {
		result.Subtitle(s.Topic)
	}

	if s.Alias != nil {
		result.Meta(*s.Alias)
	}

	for _, line := range []string{s.Description, shortenPaths(s.Files)} {
		if line != "" {
			result.Meta(line)
		}
	}

	var counts []string

	if s.TurnCount > 0 {
		counts = append(counts, fmt.Sprintf("%d turns", s.TurnCount))
	}

	if s.Lines > 0 {
		counts = append(counts, fmt.Sprintf("%d lines", s.Lines))
	}

	if s.ContextPercent > 0 {
		counts = append(counts, fmt.Sprintf("%d%% context", s.ContextPercent))
	}

	counts = append(counts, time.Relative(activity(s)))
	result.Meta(join.Middot(counts)).State(component.LabelChips(labels))

	if latest != nil {
		result.State(html.P(html.Em(gomponents.Text(latest.Body))))
	}

	return result.Node()
}
