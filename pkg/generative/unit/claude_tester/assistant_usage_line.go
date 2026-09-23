package claude_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func AssistantUsageLine(
	messageIdentifier string,
	request string,
	model string,
	usage string,
) string {
	return join.Empty(
		fmt.Sprintf(
			`{"type":"assistant","timestamp":"2026-07-20T10:00:00.000Z","requestId":%q,"message":{"id":%q,"role":"assistant","model":%q,"usage":%s}}`,
			request,
			messageIdentifier,
			model,
			usage,
		),
		"\n",
	)
}
