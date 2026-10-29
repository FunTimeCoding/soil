package fixed_generator

import "github.com/funtimecoding/soil/pkg/tool/goraidparsed/types/report"

func New(
	result *report.Report,
	failure error,
) *Generator {
	return &Generator{result: result, failure: failure}
}
