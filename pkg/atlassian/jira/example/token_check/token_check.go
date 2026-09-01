package token_check

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func TokenCheck() {
	host := environment.Required(constant.HostEnvironment)
	user := environment.Required(constant.UserEnvironment)
	token := environment.Required(constant.TokenEnvironment)
	k := environment.Required(constant.JiraDefaultProjectKeyEnvironment)
	console.Line("TokenCheck: raw /myself")
	rawMyself(host, user, token)
	console.Line()
	console.Line("TokenCheck: raw /search")
	rawSearch(host, user, token, k)
	console.Line()
	console.Line("TokenCheck: SearchLimit(1)")
	issues := common.Jira().MustSearchLimit(
		1,
		"project = %s ORDER BY updated DESC",
		k,
	)
	console.Format("  Count: %d\n", len(issues))

	for _, i := range issues {
		console.Format("  Issue: %s\n", i.Key)
	}

	console.Line("TokenCheck: done")
}
