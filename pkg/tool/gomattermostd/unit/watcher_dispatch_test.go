package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
	"time"
)

func TestWatcherDispatchReplyInWatchedThread(t *testing.T) {
	o := worker_tester.NewDispatchWatcher(t)
	o.Client.AddPost("alfa")
	o.Client.AddUser("foxtrot", "Foxtrot")
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Watcher.Index("alfa")
	o.Watcher.Dispatch(
		worker_tester.PostEvent(
			&model.Post{
				Id:       "reply",
				RootId:   "alfa",
				UserId:   "foxtrot",
				Message:  "first message",
				CreateAt: dispatchAt().UnixMilli(),
			},
		),
	)
	events := o.Watcher.Drain()["alfa"]
	assert.Count(t, 1, events)
	assert.String(t, "Foxtrot", events[0].Author)
	assert.String(t, "first message", events[0].Text)
	assert.True(t, events[0].At.Equal(dispatchAt()))
}

func TestWatcherDispatchIgnoresUnwatchedThread(t *testing.T) {
	o := worker_tester.NewDispatchWatcher(t)
	o.Client.AddUser("foxtrot", "Foxtrot")
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Watcher.Dispatch(
		worker_tester.PostEvent(
			&model.Post{
				Id:      "reply",
				RootId:  "other",
				UserId:  "foxtrot",
				Message: "not ours",
			},
		),
	)
	assert.Count(t, 0, o.Watcher.Drain())
}

func TestWatcherDispatchIgnoresOwnPost(t *testing.T) {
	o := worker_tester.NewDispatchWatcher(t)
	o.Client.AddPost("alfa")
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Watcher.Index("alfa")
	o.Watcher.Dispatch(
		worker_tester.PostEvent(
			&model.Post{
				Id:      "reply",
				RootId:  "alfa",
				UserId:  "selfuser",
				Message: "posted by the daemon",
			},
		),
	)
	assert.Count(t, 0, o.Watcher.Drain())
}

func TestWatcherDispatchReactionResolvesThroughIndex(t *testing.T) {
	o := worker_tester.NewDispatchWatcher(t)
	o.Client.AddPost("alfa")
	o.Client.AddReply("alfa", "reply")
	o.Client.AddUser("golf", "Golf")
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Watcher.Index("alfa")
	o.Watcher.Dispatch(
		worker_tester.ReactionEvent(
			model.WebsocketEventReactionAdded,
			&model.Reaction{
				UserId:    "golf",
				PostId:    "reply",
				EmojiName: "eyes",
			},
		),
	)
	events := o.Watcher.Drain()["alfa"]
	assert.Count(t, 1, events)
	assert.True(t, events[0].Kind == constant.ReactionAddedEvent)
	assert.String(t, "Golf", events[0].Author)
	assert.String(t, ":eyes:", events[0].Text)
}

func TestWatcherDispatchReactionOnUnknownPostIgnored(t *testing.T) {
	o := worker_tester.NewDispatchWatcher(t)
	o.Client.AddPost("alfa")
	o.Client.AddUser("golf", "Golf")
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Watcher.Index("alfa")
	o.Watcher.Dispatch(
		worker_tester.ReactionEvent(
			model.WebsocketEventReactionAdded,
			&model.Reaction{
				UserId:    "golf",
				PostId:    "elsewhere",
				EmojiName: "eyes",
			},
		),
	)
	assert.Count(t, 0, o.Watcher.Drain())
}

func TestWatcherForgetStopsMatching(t *testing.T) {
	o := worker_tester.NewDispatchWatcher(t)
	o.Client.AddPost("alfa")
	o.Client.AddUser("foxtrot", "Foxtrot")
	o.Store.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	o.Watcher.Index("alfa")
	o.Watcher.Forget("alfa")
	o.Watcher.Dispatch(
		worker_tester.PostEvent(
			&model.Post{
				Id:      "reply",
				RootId:  "alfa",
				UserId:  "foxtrot",
				Message: "after unsubscribe",
			},
		),
	)
	assert.Count(t, 0, o.Watcher.Drain())
}

func TestWatcherDueHoldsARootInsideItsWindow(t *testing.T) {
	o := worker_tester.NewDispatchWatcher(t)
	o.Client.AddPost("alfa")
	o.Client.AddUser("foxtrot", "Foxtrot")
	o.Watcher.Index("alfa")
	o.Watcher.Dispatch(
		worker_tester.PostEvent(
			&model.Post{
				Id:      "reply",
				RootId:  "alfa",
				UserId:  "foxtrot",
				Message: "first message",
			},
		),
	)
	assert.Count(t, 0, o.Watcher.Due(time.Now(), constant.DebounceWindow))
	assert.Count(
		t,
		1,
		o.Watcher.Due(
			time.Now().Add(2*constant.DebounceWindow),
			constant.DebounceWindow,
		),
	)
	assert.Count(t, 0, o.Watcher.Drain())
}
