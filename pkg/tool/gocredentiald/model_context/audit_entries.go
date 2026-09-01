package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) auditEntries(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Audit,
) (*mcp.CallToolResult, error) {
	years := int(a.StaleYears)

	if years == 0 {
		years = constant.DefaultStaleYears
	}

	report := s.service.Audit(years)

	if a.Bucket == "" {
		return response.SuccessAny(
			map[string]any{
				constant.BucketStale:         auditSample(report.Stale),
				constant.BucketEmptyUser:     auditSample(report.EmptyUser),
				constant.BucketEmptyPassword: auditSample(report.EmptyPassword),
				constant.BucketDuplicates:    auditSample(report.Duplicates),
			},
		)
	}

	bucket := auditBucket(report, a.Bucket)

	if bucket == nil {
		return response.Fail("unknown bucket: %s", a.Bucket)
	}

	limit := int(a.Limit)

	if limit == 0 {
		limit = constant.AuditPageLimit
	}

	offset := int(a.Offset)

	if offset > len(bucket) {
		offset = len(bucket)
	}

	end := offset + limit

	if end > len(bucket) {
		end = len(bucket)
	}

	return response.SuccessAny(
		map[string]any{
			"bucket": a.Bucket,
			"total":  len(bucket),
			"offset": offset,
			"rows":   describeCredentials(bucket[offset:end]),
		},
	)
}
