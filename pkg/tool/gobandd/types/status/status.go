package status

import "github.com/funtimecoding/soil/pkg/tool/gobandd/types/operation"

type Status struct {
	Instance      string                 `json:"instance"`
	Host          string                 `json:"host"`
	Provisioning  string                 `json:"provisioning"`
	ScreenEnabled bool                   `json:"screen_enabled"`
	ScreenTimeout int                    `json:"screen_timeout"`
	Listener      bool                   `json:"listener"`
	Consent       string                 `json:"consent"`
	Power         string                 `json:"power"`
	Available     []string               `json:"available"`
	Operations    []*operation.Operation `json:"operations"`
}
