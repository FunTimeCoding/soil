package unit

import (
	"github.com/funtimecoding/soil/pkg/crap"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/crap/report"
)

func build(coverage map[string]float64) *report.Report {
	i := index.New("/r")
	i.Add(function.New("m", "/r/a.go", 1, "A", 2))
	i.Add(function.New("m", "/r/b.go", 1, "B", 6))

	return crap.Build(i, coverage, constant.Pessimistic)
}
