package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/decision_page"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func pushDecisions(
	w http.ResponseWriter,
	decisions []*decision.Decision,
) {
	layout.PushEvent(w, constant.DecisionEvent, session_page.Grid(decisions))

	for _, d := range decisions {
		layout.PushEvent(
			w,
			decision_page.Event(d.Identifier, constant.StatusPart),
			decision_page.Status(d),
		)
		layout.PushEvent(
			w,
			decision_page.Event(d.Identifier, constant.ChoicePart),
			decision_page.Choices(d),
		)
		layout.PushEvent(
			w,
			decision_page.Event(d.Identifier, constant.DismissPart),
			decision_page.Dismissals(d),
		)
		layout.PushEvent(
			w,
			decision_page.Event(d.Identifier, constant.ThreadPart),
			decision_page.Thread(d),
		)
	}
}
