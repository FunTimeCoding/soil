package bidi_server

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/message"
)

func (s *Server) Answer(
	method string,
	result string,
) {
	r := message.NewReply()
	r.Type = constant.SuccessType
	r.Result = json.RawMessage(result)
	s.mutex.Lock()
	s.answers[method] = r
	s.mutex.Unlock()
}
