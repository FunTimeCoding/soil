package service

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

type Service struct {
	store    *store.Store
	ollama   *ollama.Client
	reranker *rerank.Reranker
}
