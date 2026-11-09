package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/tokenizer"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	queryFace "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"time"
)

type Service struct {
	reporter      face.Reporter
	store         *store.Store
	query         queryFace.Client
	guard         *peer.Guard
	tokenizer     *tokenizer.Encoder
	hiddenTag     string
	referenceRoot string
	clock         func() time.Time
}
