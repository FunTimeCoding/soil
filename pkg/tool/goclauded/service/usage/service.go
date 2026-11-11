package usage

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
	"github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

type Service struct {
	store            *store.Store
	notifier         soilFace.EventNotifier
	browser          *chromium.Client
	utilization      face.AnthropicClient
	utilizationGuard *peer.Guard
	credential       func() *credential.Credential
	usageUnreadable  bool
	clock            func() time.Time
	logger           *logger.Logger
	reporter         soilFace.Reporter
}
