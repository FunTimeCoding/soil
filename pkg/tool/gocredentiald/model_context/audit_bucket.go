package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/audit_report"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"
)

func auditBucket(
	report *audit_report.Report,
	name string,
) []*credential.Credential {
	switch name {
	case constant.BucketStale:
		return report.Stale
	case constant.BucketEmptyUser:
		return report.EmptyUser
	case constant.BucketEmptyPassword:
		return report.EmptyPassword
	case constant.BucketDuplicates:
		return report.Duplicates
	}

	return nil
}
