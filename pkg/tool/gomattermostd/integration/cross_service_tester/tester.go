package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
)

type Tester struct {
	Upstream  *mattermost_client_tester.Tester
	Store     *store.Store
	Watcher   *watcher.Watcher
	Goclauded *base.Server
	Reporter  *memory.Memory
}
