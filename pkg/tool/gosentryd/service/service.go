package service

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	client       face.SentryClient
	guard        *peer.Guard
	organization string
}
