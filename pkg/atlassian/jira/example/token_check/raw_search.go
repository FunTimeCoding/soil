package token_check

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func rawSearch(
	host string,
	user string,
	token string,
	project string,
) {
	rawGet(
		locator.New(host).Path("rest/api/2/search/jql").Set(
			constant.JiraQueryKey,
			fmt.Sprintf("project = %s ORDER BY updated DESC", project),
		).Set(constant.JiraMaximumResultsKey, "1").String(),
		user,
		token,
	)
}
