package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
)

func RemediationItem(v *remediation_item.Item) server.RemediationItem {
	result := server.RemediationItem{
		Subject:         v.Subject,
		Kind:            v.Kind,
		Release:         optional(v.Release),
		Source:          optional(v.Source),
		Reason:          optional(v.Reason),
		Hosts:           v.Hosts,
		Vulnerabilities: v.Vulnerabilities,
		TopScore:        v.TopScore,
		Unknown:         v.Unknown,
		Addressed:       v.Addressed,
	}

	if result.Vulnerabilities == nil {
		result.Vulnerabilities = []string{}
	}

	if v.Count != nil {
		c := AffectedCount(v.Count)
		result.Count = &c
	}

	if v.Decision != nil {
		result.Decision = Flag(v.Decision)
	}

	return result
}
