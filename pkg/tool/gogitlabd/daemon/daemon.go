package daemon

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web"
	gitlabWorker "github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/funtimecoding/soil/pkg/worker"
	"gorm.io/gorm"
	"sync"
)

type Daemon struct {
	inventory   *instance.Inventory
	pool        *pool.Pool[instance.Instance, *instance.Instance, *peered.Peered]
	selection   *selection.Selection
	service     *service.Service
	web         *web.Server
	worker      *gitlabWorker.Worker
	database    *gorm.DB
	connector   face.Client
	rosterGuard *peer.Guard
	clocks      []*worker.Worker
	closeOnce   sync.Once
}
