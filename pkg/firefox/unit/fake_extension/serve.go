package fake_extension

import (
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"testing"
)

func (x *Extension) serve(t *testing.T) {
	for {
		var r message.Request

		if wsjson.Read(t.Context(), x.connection, &r) != nil {
			return
		}

		x.mutex.Lock()
		reply, okay := x.replies[r.Method]
		x.mutex.Unlock()

		if !okay {
			reply = message.NewFailedReply(
				join.Empty("unknown method: ", r.Method),
			)
		}

		answer := *reply
		answer.Identifier = r.Identifier

		if wsjson.Write(t.Context(), x.connection, answer) != nil {
			return
		}
	}
}
