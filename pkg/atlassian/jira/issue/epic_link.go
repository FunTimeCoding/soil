package issue

import "github.com/funtimecoding/soil/pkg/atlassian/jira/constant"

func (i *Issue) EpicLink() string {
	if i.Type != EpicType {
		if l := i.CustomValue(constant.ParentEpic); ValidValue(l) {
			return l
		}
	}

	return ""
}
