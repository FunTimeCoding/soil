package token_check

import "github.com/funtimecoding/soil/pkg/web/locator"

func rawMyself(
	host string,
	user string,
	token string,
) {
	rawGet(
		locator.New(host).Path("rest/api/2/myself").String(),
		user,
		token,
	)
}
