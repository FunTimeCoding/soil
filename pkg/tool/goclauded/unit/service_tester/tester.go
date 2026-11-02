package service_tester

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/document"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/startup"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/double"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"sync"
	"testing"
)

type Tester struct {
	*double.Set
	Store          *store_tester.Tester
	t              *testing.T
	Roster         *roster.Service
	Session        *session.Service
	Startup        *startup.Service
	Label          *label.Service
	Message        *message.Service
	Memory         *memory.Service
	Event          *event.Service
	Transcript     *transcript.Service
	Usage          *usage.Service
	Index          *search_index.Index
	Search         *search.Service
	Document       *document.Service
	searchDatabase *sql.DB
	memoryGuard    *peer.Guard
	queryGuard     *peer.Guard
	stop           *stopper.Stopper
	closeOnce      sync.Once
}
