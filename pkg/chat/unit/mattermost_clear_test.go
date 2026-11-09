package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/mattermost/mattermost/server/public/model"
	"sync"
	"testing"
	"time"
)

func TestClearDeletesRootsAndReplies(t *testing.T) {
	var deleted []string
	var guard sync.Mutex
	c := mattermost_client_tester.New(
		t,
		threadChannel(
			&deleted,
			&guard,
			&model.Post{Id: "root", UserId: "bot", CreateAt: 1},
			&model.Post{
				Id:       "reply",
				RootId:   "root",
				UserId:   "alfa",
				CreateAt: 2,
			},
		),
	).Client
	n, e := c.Clear(&model.Channel{Id: "echo"})
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, n)
	guard.Lock()
	defer guard.Unlock()
	assert.Any(t, []string{"reply", "root"}, deleted)
}

func TestClearOldKeepsNewestAndRecent(t *testing.T) {
	var deleted []string
	var guard sync.Mutex
	cutoff := time.UnixMilli(100)
	c := mattermost_client_tester.New(
		t,
		threadChannel(
			&deleted,
			&guard,
			&model.Post{Id: "oldest", UserId: "bot", CreateAt: 10},
			&model.Post{Id: "older", UserId: "bot", CreateAt: 20},
			&model.Post{Id: "recent", UserId: "bot", CreateAt: 200},
			&model.Post{Id: "newest", UserId: "bot", CreateAt: 300},
		),
	).Client
	n, skipped, e := c.ClearOld(&model.Channel{Id: "echo"}, cutoff, 3)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, n)
	assert.Integer(t, 0, skipped)
	guard.Lock()
	defer guard.Unlock()
	assert.Any(t, []string{"older", "oldest"}, deleted)
}
