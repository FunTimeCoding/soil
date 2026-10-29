package fixed_generator

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/types/report"
	"time"
)

func (g *Generator) Generate(
	[]string,
	time.Time,
) (*report.Report, error) {
	return g.result, g.failure
}
