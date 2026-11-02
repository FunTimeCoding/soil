package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	queryFace "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

func New(
	s *store.Store,
	q queryFace.Client,
	g *peer.Guard,
	hiddenTag string,
	referenceRoot string,
	r face.Reporter,
) *Service {
	return &Service{
		reporter:      r,
		store:         s,
		query:         q,
		guard:         g,
		hiddenTag:     hiddenTag,
		referenceRoot: referenceRoot,
	}
}
