package model_context

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"

func (s *Server) registerResources() {
	s.guide(
		constant.MatchingGuideLocator,
		"Matching: sources, coverage, clean, fix kinds, the kernel",
		"How installed packages become findings: the National Vulnerability Database and the Debian and Freexian trackers, coverage and the clean rule, the fix kinds, kernel matching and states, the score filter.",
		constant.MatchingGuide,
	)
	s.guide(
		constant.RemediationGuideLocator,
		"Remediation: addressed, items, subjects, decisions",
		"What addressed means, how each finding becomes one remediation item, the subjects decisions are keyed by, and how to record a mitigation or accepted risk.",
		constant.RemediationGuide,
	)
}
