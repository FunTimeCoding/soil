package source

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"time"
)

type Source struct {
	BaseWeight      float64
	AgeMultiplier   float64
	StaleThreshold  time.Duration
	StaleMultiplier float64
	SeverityWeights map[soil.Severity]float64
	TriageBoost     float64
}
