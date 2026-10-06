package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
)

func joinedConcerns(r *output.Results) string {
	return strings.Join(service_tester.ConcernText(r), "\n")
}
