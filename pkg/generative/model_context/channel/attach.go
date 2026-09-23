package channel

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Server) Attach() string {
	<-s.ready
	meta := map[string]string{
		constant.ChannelKindMeta:  constant.ChannelAttachKind,
		constant.ChannelNonceMeta: s.nonce,
	}
	wait := constant.ChannelAttachInterval

	for !s.Opened() {
		s.Push(constant.ChannelAttachMessage, meta)
		s.sleep(wait)
		wait *= 2

		if wait > constant.ChannelAttachMaximum {
			wait = constant.ChannelAttachMaximum
		}
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.callsign
}
