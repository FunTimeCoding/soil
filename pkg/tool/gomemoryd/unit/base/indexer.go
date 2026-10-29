package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"

func (s *Server) Indexer() *mock_client.Client {
	return s.indexer
}
