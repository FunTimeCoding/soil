package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service/format"
)

func ProfileText(result *ProfileResult) string {
	var section []string

	if len(result.Always) > 0 {
		body := make([]string, 0, len(result.Always))

		for i := range result.Always {
			body = append(body, format.AlwaysMemory(&result.Always[i]))
		}

		section = append(
			section,
			sectionText(constant.AlwaysSectionHeading, body),
		)
	}

	if len(result.Index) > 0 {
		body := make([]string, 0, len(result.Index))

		for i := range result.Index {
			body = append(body, format.IndexEntry(&result.Index[i]))
		}

		section = append(
			section,
			sectionText(constant.IndexSectionHeading, body),
		)
	}

	if len(result.Relevant) > 0 {
		body := make([]string, 0, len(result.Relevant))

		for i := range result.Relevant {
			body = append(body, format.RelevantMemory(&result.Relevant[i]))
		}

		section = append(
			section,
			sectionText(constant.RelevantSectionHeading, body),
		)
	}

	if len(result.Impressions) > 0 {
		body := make([]string, 0, len(result.Impressions))

		for i := range result.Impressions {
			body = append(body, format.Impression(&result.Impressions[i]))
		}

		section = append(
			section,
			sectionText(constant.ImpressionSectionHeading, body),
		)
	}

	if len(result.Completions) > 0 {
		body := make([]string, 0, len(result.Completions))

		for _, c := range result.Completions {
			body = append(body, format.Completion(c.SessionName, c.Body))
		}

		section = append(
			section,
			sectionText(constant.CompletionSectionHeading, body),
		)
	}

	return join.NewLine(section)
}
