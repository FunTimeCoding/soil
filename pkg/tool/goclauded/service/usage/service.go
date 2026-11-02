package usage

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
	"github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

type Service struct {
	store            *store.Store
	notifier         goclauded.Notifier
	browser          *chromium.Client
	utilization      face.AnthropicClient
	utilizationGuard *peer.Guard
	credential       func() *credential.Credential
	usageUnreadable  bool
	clock            func() time.Time
	logger           *logger.Logger
	reporter         library.Reporter
}
