package gmail

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

func (c *Client) profile(s *gmail.Service) *gmail.Profile {
	result, e := s.Users.GetProfile("me").Do()
	errors.PanicOnError(e)

	return result
}
