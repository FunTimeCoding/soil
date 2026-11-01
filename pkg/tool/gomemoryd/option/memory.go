package option

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"

type Memory struct {
	Address       string
	ServiceTokens []string
	LitePath      string
	HiddenTag     string
	ReferenceRoot string
	Query         func() face.Client
}
