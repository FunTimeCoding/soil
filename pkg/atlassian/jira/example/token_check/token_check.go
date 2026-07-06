package token_check

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	jiraConstant "github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func TokenCheck() {
	host := environment.Required(constant.HostEnvironment)
	user := environment.Required(constant.UserEnvironment)
	token := environment.Required(constant.TokenEnvironment)
	k := environment.Required(jiraConstant.DefaultProjectKeyEnvironment)
	fmt.Println("TokenCheck: raw /myself")
	rawMyself(host, user, token)
	fmt.Println()
	fmt.Println("TokenCheck: raw /search")
	rawSearch(host, user, token, k)
	fmt.Println()
	fmt.Println("TokenCheck: SearchLimit(1)")
	issues := common.Jira().MustSearchLimit(
		1,
		"project = %s ORDER BY updated DESC",
		k,
	)
	fmt.Printf("  Count: %d\n", len(issues))

	for _, i := range issues {
		fmt.Printf("  Issue: %s\n", i.Key)
	}

	fmt.Println("TokenCheck: done")
}
