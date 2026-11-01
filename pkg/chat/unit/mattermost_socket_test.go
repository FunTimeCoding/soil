package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/socket"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
)

func TestASocketThatCannotRefreshIsNotOpened(t *testing.T) {
	m := mock_client.New("alfa")
	defer m.Close()
	m.Failure = unavailable.Format("mattermost restarting")
	s := socket.New(m, func(*model.WebSocketEvent) {})
	defer s.Close()
	assert.True(t, unavailable.Is(s.Open(context.Background())))
}

func TestClosingASocketThatNeverOpenedIsQuiet(t *testing.T) {
	m := mock_client.New("alfa")
	defer m.Close()
	socket.New(m, func(*model.WebSocketEvent) {}).Close()
}
