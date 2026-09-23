package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
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
	o.Watcher.Flush("alfa")
	result := o.Sink.All()
	assert.Integer(t, 1, len(result))
	assert.String(t, "kilo", result[0].Callsign)
	assert.String(t, "[papa] 14:32 · Foxtrot: first message", result[0].Body)
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
	o.Watcher.Flush("other")
	assert.Integer(t, 0, len(o.Sink.All()))
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
	o.Watcher.Flush("alfa")
	assert.Integer(t, 0, len(o.Sink.All()))
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
	o.Watcher.Flush("alfa")
	result := o.Sink.All()
	assert.Integer(t, 1, len(result))
	assert.StringContains(t, "Golf reacted :eyes:", result[0].Body)
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
	o.Watcher.Flush("alfa")
	assert.Integer(t, 0, len(o.Sink.All()))
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
	o.Watcher.Flush("alfa")
	assert.Integer(t, 0, len(o.Sink.All()))
}
