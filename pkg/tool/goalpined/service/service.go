package service

import "github.com/funtimecoding/soil/pkg/tool/goalpined/face"

type Service struct {
	packageRoot  string
	keyDirectory string
	signatureKey string
	indexer      face.Indexer
}
