package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"

func (s *Server) Indexer() *mock_indexer.Indexer {
	return s.indexer
}
