package base

import (
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_indexer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
)

type Server struct {
	Upstream      *mattermost_client_tester.Tester
	Store         *store.Store
	Indexer       *mock_indexer.Indexer
	ContextServer *model_context_server.Server
}
