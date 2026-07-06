package service

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/tokenizer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

type Service struct {
	store     *store.Store
	indexer   face.Indexer
	searcher  face.Searcher
	lister    face.Lister
	tokenizer *tokenizer.Encoder
}
