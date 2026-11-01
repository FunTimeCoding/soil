package service

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/tokenizer"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"sync"
)

type Service struct {
	logger        *logger.Logger
	database      *sql.DB
	store         *store.Store
	query         face.Client
	tokenizer     *tokenizer.Encoder
	hiddenTag     string
	referenceRoot string
	closeOnce     sync.Once
}
