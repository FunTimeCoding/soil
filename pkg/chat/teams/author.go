package teams

import "github.com/funtimecoding/soil/pkg/chat/teams/response"

func author(s *response.Sender) string {
	switch {
	case s == nil:
		return ""
	case s.User != nil:
		return s.User.DisplayName
	case s.Application != nil:
		return s.Application.DisplayName
	default:
		return ""
	}
}
