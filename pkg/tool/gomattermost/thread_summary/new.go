package thread_summary

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"

func New(
	thread client.Thread,
	summary string,
) *Summary {
	return &Summary{Thread: thread, Summary: summary}
}
