package face

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/types/report"
	"time"
)

type Generator interface {
	Generate(
		files []string,
		date time.Time,
	) (*report.Report, error)
}
