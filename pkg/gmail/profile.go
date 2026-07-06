package gmail

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

func (c *Client) profile(s *gmail.Service) *gmail.Profile {
	// Pass service because it is not initialized yet
	result, e := s.Users.GetProfile("me").Do()
	errors.PanicOnError(e)

	return result
}
