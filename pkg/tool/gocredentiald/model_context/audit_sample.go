package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"
)

func auditSample(bucket []*credential.Credential) map[string]any {
	sample := bucket

	if len(sample) > constant.AuditSampleLimit {
		sample = sample[:constant.AuditSampleLimit]
	}

	return map[string]any{
		"total":  len(bucket),
		"sample": describeCredentials(sample),
	}
}
