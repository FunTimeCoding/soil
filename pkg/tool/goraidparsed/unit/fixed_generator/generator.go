package fixed_generator

import "github.com/funtimecoding/soil/pkg/tool/goraidparsed/types/report"

type Generator struct {
	result  *report.Report
	failure error
}
