package option

import "github.com/funtimecoding/soil/pkg/tool/goalpined/face"

type Alpine struct {
	Address       string
	ServiceTokens []string
	PackageRoot   string
	KeyDirectory  string
	SignatureKey  string
	Indexer       func() face.Indexer
}
