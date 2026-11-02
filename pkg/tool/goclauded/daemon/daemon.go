package daemon

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/face"
	generative "github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/gauge"
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
	"github.com/funtimecoding/soil/pkg/tool/goclauded/watcher"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	database         *gorm.DB
	searchDatabase   *sql.DB
	memoryPeer       memoryd.Client
	memoryGuard      *peer.Guard
	query            queryd.Client
	queryGuard       *peer.Guard
	browser          *chromium.Client
	utilization      generative.AnthropicClient
	utilizationGuard *peer.Guard
	stop             *stopper.Stopper
	roster           *roster.Service
	session          *session.Service
	startup          *startup.Service
	label            *label.Service
	message          *message.Service
	memory           *memory.Service
	event            *event.Service
	transcript       *transcript.Service
	usage            *usage.Service
	search           *search.Service
	document         *document.Service
	web              *web.Server
	gauge            *gauge.Gauge
	watcher          *watcher.Watcher
	clocks           []face.Worker
	logger           *logger.Logger
	closeOnce        sync.Once
}
