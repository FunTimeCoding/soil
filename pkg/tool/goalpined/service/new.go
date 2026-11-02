package service

import "github.com/funtimecoding/soil/pkg/tool/goalpined/face"

func New(
	packageRoot string,
	keyDirectory string,
	signatureKey string,
	i face.Indexer,
) *Service {
	return &Service{
		packageRoot:  packageRoot,
		keyDirectory: keyDirectory,
		signatureKey: signatureKey,
		indexer:      i,
	}
}
