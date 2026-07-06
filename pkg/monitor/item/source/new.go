package source

import (
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"time"
)

func New(
	baseWeight float64,
	ageMultiplier float64,
	staleThreshold time.Duration,
	staleMultiplier float64,
	triageBoost float64,
	severityWeights map[constant.Severity]float64,
) *Source {
	if severityWeights == nil {
		panic("nil severity weights")
	}

	return &Source{
		BaseWeight:      baseWeight,
		AgeMultiplier:   ageMultiplier,
		StaleThreshold:  staleThreshold,
		StaleMultiplier: staleMultiplier,
		SeverityWeights: severityWeights,
		TriageBoost:     triageBoost,
	}
}
