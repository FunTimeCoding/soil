package source

import (
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"time"
)

type Source struct {
	BaseWeight      float64
	AgeMultiplier   float64
	StaleThreshold  time.Duration
	StaleMultiplier float64
	SeverityWeights map[constant.Severity]float64
	TriageBoost     float64
}
