package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"github.com/mattermost/mattermost/server/public/model"
	"testing"
	"time"
)

func newDispatchWatcher(
	t *testing.T,
) (*watcher.Watcher, *store.Store, *mock_client.Client, *notifySink) {
	t.Helper()
	sink, c := newNotifyClient(t)
	s := store.New(lite.NewMemory())
	r := memory.New()
	chat := mock_client.New("selfuser")

	return watcher.New(
		chat,
		s,
		notifier.New(c, "mattermost", r),
		logger.New(context.Background()),
		r,
		time.Hour,
	), s, chat, sink
}

func dispatchAt() time.Time {
	return time.Date(2026, 9, 8, 14, 32, 0, 0, time.Local)
}

func postEvent(p *model.Post) *model.WebSocketEvent {
	v := &model.WebSocketEvent{}
	v = v.SetEvent(model.WebsocketEventPosted)

	return v.SetData(
		map[string]any{constant.MattermostPostField: notation.Encode(p, false)},
	)
}

func reactionEvent(
	kind model.WebsocketEventType,
	r *model.Reaction,
) *model.WebSocketEvent {
	v := &model.WebSocketEvent{}
	v = v.SetEvent(kind)

	return v.SetData(
		map[string]any{
			constant.MattermostReactionField: notation.Encode(r, false),
		},
	)
}

func TestWatcherDispatchReplyInWatchedThread(t *testing.T) {
	w, s, chat, sink := newDispatchWatcher(t)
	chat.AddPost("alfa")
	chat.AddUser("foxtrot", "Foxtrot")
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Index("alfa")
	w.Dispatch(
		postEvent(
			&model.Post{
				Id:       "reply",
				RootId:   "alfa",
				UserId:   "foxtrot",
				Message:  "first message",
				CreateAt: dispatchAt().UnixMilli(),
			},
		),
	)
	w.Flush("alfa")
	result := sink.all()
	assert.Integer(t, 1, len(result))
	assert.String(t, "kilo", result[0].Callsign)
	assert.String(t, "[papa] 14:32 · Foxtrot: first message", result[0].Body)
}

func TestWatcherDispatchIgnoresUnwatchedThread(t *testing.T) {
	w, s, chat, sink := newDispatchWatcher(t)
	chat.AddUser("foxtrot", "Foxtrot")
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Dispatch(
		postEvent(
			&model.Post{
				Id:      "reply",
				RootId:  "other",
				UserId:  "foxtrot",
				Message: "not ours",
			},
		),
	)
	w.Flush("other")
	assert.Integer(t, 0, len(sink.all()))
}

func TestWatcherDispatchIgnoresOwnPost(t *testing.T) {
	w, s, chat, sink := newDispatchWatcher(t)
	chat.AddPost("alfa")
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Index("alfa")
	w.Dispatch(
		postEvent(
			&model.Post{
				Id:      "reply",
				RootId:  "alfa",
				UserId:  "selfuser",
				Message: "posted by the daemon",
			},
		),
	)
	w.Flush("alfa")
	assert.Integer(t, 0, len(sink.all()))
}

func TestWatcherDispatchReactionResolvesThroughIndex(t *testing.T) {
	w, s, chat, sink := newDispatchWatcher(t)
	chat.AddPost("alfa")
	chat.AddReply("alfa", "reply")
	chat.AddUser("golf", "Golf")
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Index("alfa")
	w.Dispatch(
		reactionEvent(
			model.WebsocketEventReactionAdded,
			&model.Reaction{
				UserId:    "golf",
				PostId:    "reply",
				EmojiName: "eyes",
			},
		),
	)
	w.Flush("alfa")
	result := sink.all()
	assert.Integer(t, 1, len(result))
	assert.StringContains(t, "Golf reacted :eyes:", result[0].Body)
}

func TestWatcherDispatchReactionOnUnknownPostIgnored(t *testing.T) {
	w, s, chat, sink := newDispatchWatcher(t)
	chat.AddPost("alfa")
	chat.AddUser("golf", "Golf")
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Index("alfa")
	w.Dispatch(
		reactionEvent(
			model.WebsocketEventReactionAdded,
			&model.Reaction{
				UserId:    "golf",
				PostId:    "elsewhere",
				EmojiName: "eyes",
			},
		),
	)
	w.Flush("alfa")
	assert.Integer(t, 0, len(sink.all()))
}

func TestWatcherForgetStopsMatching(t *testing.T) {
	w, s, chat, sink := newDispatchWatcher(t)
	chat.AddPost("alfa")
	chat.AddUser("foxtrot", "Foxtrot")
	s.MustCreate(subscription.New("kilo", "alfa", "bravo", "papa"))
	w.Index("alfa")
	w.Forget("alfa")
	w.Dispatch(
		postEvent(
			&model.Post{
				Id:      "reply",
				RootId:  "alfa",
				UserId:  "foxtrot",
				Message: "after unsubscribe",
			},
		),
	)
	w.Flush("alfa")
	assert.Integer(t, 0, len(sink.all()))
}
