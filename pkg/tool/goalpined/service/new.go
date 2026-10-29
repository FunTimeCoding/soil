package service

import "github.com/funtimecoding/soil/pkg/tool/goalpined/option"

func New(o *option.Alpine) *Service {
	return &Service{
		packageRoot:  o.PackageRoot,
		keyDirectory: o.KeyDirectory,
		signatureKey: o.SignatureKey,
		indexer:      o.Indexer(),
	}
}
