package result

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

type Failed struct {
	Report  *client.FailedReport `json:"report"`
	Summary *FailedSummary       `json:"summary"`
}
