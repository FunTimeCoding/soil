package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost_client_tester"
	"github.com/mattermost/mattermost/server/public/model"
	"sync"
	"testing"
)

func TestWipePrefixTakesPrefixedReply(t *testing.T) {
	var deleted []string
	var guard sync.Mutex
	c := mattermost_client_tester.New(
		t,
		mattermost_client_tester.ThreadChannel(
			&deleted,
			&guard,
			&model.Post{
				Id:       "root",
				UserId:   "bot",
				Message:  "wipe: root",
				CreateAt: 1,
			},
			&model.Post{
				Id:       "reply",
				RootId:   "root",
				UserId:   "bot",
				Message:  "wipe: reply",
				CreateAt: 2,
			},
			&model.Post{
				Id:       "other",
				UserId:   "alfa",
				Message:  "wipe: hers",
				CreateAt: 3,
			},
		),
	).Client
	c.WipePrefix(&model.Channel{Id: "echo"}, "wipe:")
	guard.Lock()
	defer guard.Unlock()
	assert.Any(t, []string{"root", "reply"}, deleted)
}
