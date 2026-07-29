package issue

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func (i *Issue) EpicLink() string {
	if i.Type != constant.JiraEpicType {
		if l := i.CustomValue(constant.JiraParentEpic); ValidValue(l) {
			return l
		}
	}

	return ""
}
