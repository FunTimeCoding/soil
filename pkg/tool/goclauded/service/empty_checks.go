package service

func (s *Service) emptyChecks() []emptyCheck {
	return []emptyCheck{
		{s.store.CountSessionCompletions, "session has completions"},
		{s.store.CountSessionSummaries, "session has a summary"},
		{s.store.CountSessionLabels, "session has labels"},
		{s.store.CountSessionPulses, "session has pulses"},
		{s.store.CountSessionContextLoads, "session has context loads"},
		{
			func(i string) (int64, error) {
				return s.store.CountSessionEventsExcluding(i, lifecycleKinds())
			},
			"session has events beyond its lifecycle",
		},
	}
}
