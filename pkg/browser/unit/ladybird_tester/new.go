package ladybird_tester

import (
	"github.com/funtimecoding/soil/pkg/browser/ladybird"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/ladybird_option"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/browser/unit/bidi_server"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	result := &Tester{
		t:      t,
		Server: bidi_server.New(t),
		events: make(chan *message.Reply, 16),
	}
	o := ladybird_option.New()
	o.Address = result.Server.Address()
	o.Event = func(r *message.Reply) { result.events <- r }
	result.Client = ladybird.New(o)
	result.Client.Open()
	t.Cleanup(result.Client.Close)

	return result
}
