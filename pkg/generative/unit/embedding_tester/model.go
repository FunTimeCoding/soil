package embedding_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
)

func Model(s *upstream_server.Server) string {
	var result sent
	notation.MustDecode(s.LastBody(), &result, false)

	return result.Model
}
