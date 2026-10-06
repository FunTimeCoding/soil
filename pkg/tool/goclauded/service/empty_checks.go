package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/deletion"
)

func (s *Service) emptyChecks() []deletion.EmptyCheck {
	return []deletion.EmptyCheck{
		{
			Count:   s.store.CountSessionCompletions,
			Message: "session has completions",
		},
		{
			Count:   s.store.CountSessionSummaries,
			Message: "session has a summary",
		},
		{Count: s.store.CountSessionLabels, Message: "session has labels"},
		{Count: s.store.CountSessionPulses, Message: "session has pulses"},
		{
			Count:   s.store.CountSessionContextLoads,
			Message: "session has context loads",
		},
		{
			Count: func(i string) (int64, error) {
				return s.store.CountSessionEventsExcluding(
					i,
					constant.LifecycleKinds,
				)
			},
			Message: "session has events beyond its lifecycle",
		},
	}
}
