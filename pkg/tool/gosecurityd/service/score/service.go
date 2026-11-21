package score

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/mirror"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/score"
	"github.com/funtimecoding/soil/pkg/vulnerability/face"
	"sync"
	"time"
)

type Service struct {
	store      *store.Store
	client     face.NationalDatabase
	guard      *peer.Guard
	clock      func() time.Time
	scores     map[string]*score.Score
	state      *mirror.Mirror
	loaded     bool
	generation uint64
	mutex      sync.RWMutex
}
