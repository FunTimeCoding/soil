package chart

import (
	"github.com/funtimecoding/soil/pkg/web/chart/series"
	"time"
)

type Chart struct {
	start      time.Time
	end        time.Time
	now        time.Time
	maximum    float64
	series     []*series.Series
	guide      bool
	projection bool
}
