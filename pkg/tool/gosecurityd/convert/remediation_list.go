package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_result"
)

func RemediationList(v *remediation_result.Result) *server.RemediationList {
	result := &server.RemediationList{
		Items:       make([]server.RemediationItem, len(v.Items)),
		Open:        v.Open,
		Addressed:   v.Addressed,
		Unimportant: v.Unimportant,
	}

	for i, item := range v.Items {
		result.Items[i] = RemediationItem(item)
	}

	return result
}
