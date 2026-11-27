package ladybird_tester

import (
	"github.com/funtimecoding/soil/pkg/browser/ladybird"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/browser/unit/bidi_server"
	"testing"
)

type Tester struct {
	t      *testing.T
	Server *bidi_server.Server
	Client *ladybird.Client
	events chan *message.Reply
}
