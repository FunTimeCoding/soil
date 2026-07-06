package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"

func (s *Server) Reranker() *rerank.Reranker {
	return s.reranker
}
