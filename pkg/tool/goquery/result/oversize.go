package result

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

type Oversize struct {
	Report  *client.OversizeReport `json:"report"`
	Summary *OversizeSummary       `json:"summary"`
}
