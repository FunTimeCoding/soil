package service

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/tokenizer"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	queryFace "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"sync"
)

type Service struct {
	reporter      face.Reporter
	database      *sql.DB
	store         *store.Store
	query         queryFace.Client
	guard         *peer.Guard
	tokenizer     *tokenizer.Encoder
	hiddenTag     string
	referenceRoot string
	closeOnce     sync.Once
}
