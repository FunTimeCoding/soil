package thread_summary

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"

type Summary struct {
	Thread  client.Thread `json:"thread"`
	Summary string        `json:"summary"`
}
