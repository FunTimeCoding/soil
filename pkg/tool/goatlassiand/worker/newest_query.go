package worker

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/query"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
)

func NewestQuery(
	project []string,
	closed string,
) string {
	return fmt.Sprintf(
		constant.NewestQuery,
		join.Comma(query.Quote(project)),
		closed,
	)
}
