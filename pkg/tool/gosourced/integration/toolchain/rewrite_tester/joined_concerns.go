package rewrite_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
)

func JoinedConcerns(r *output.Results) string {
	return strings.Join(service_tester.ConcernText(r), "\n")
}
