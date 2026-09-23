package spacing

import "github.com/funtimecoding/soil/pkg/lint/file_report"

type Spacing struct {
	report                     *file_report.Report
	path                       string
	pastLine                   string
	pastWasBlank               bool
	needBlankAfterClosingBrace bool
	inBacktick                 bool
	blockStack                 []bool
	pendingControl             bool
	parenDepth                 int
	pendingBlank               bool
	pendingBlankLine           int
}
