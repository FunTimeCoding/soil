package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"testing"
)

func TestGuidesAreServedAndReferenced(t *testing.T) {
	s := base.New(t)
	s.Started()
	tools := s.Tools()
	served := []string{}

	for _, r := range tools.ListResources() {
		served = append(served, r.URI)
	}

	assert.Strings(
		t,
		[]string{
			"gosecurityd://guide/matching",
			"gosecurityd://guide/remediation",
		},
		served,
	)
	remediation := tools.ReadResource(constant.RemediationGuideLocator).Contents[0]
	assert.String(t, "text/markdown", remediation.MIMEType)
	assert.StringContains(t, "upgrade:<release>:<source>", remediation.Text)
	assert.StringContains(
		t,
		"## Coverage and clean",
		tools.ReadResource(constant.MatchingGuideLocator).Contents[0].Text,
	)
	instructions := tools.InitializeResult().Instructions
	assert.StringContains(t, "gosecurityd://guide/matching", instructions)
	assert.StringContains(t, "gosecurityd://guide/remediation", instructions)
}

func TestToolDescriptionsStayShortAndInFullWords(t *testing.T) {
	s := base.New(t)
	s.Started()
	tools := s.Tools()
	texts := []string{tools.InitializeResult().Instructions}

	for _, v := range tools.ListTools() {
		assert.True(t, len(v.Description) <= 300)
		texts = append(texts, v.Description)
	}

	for _, text := range texts {
		for _, acronym := range []string{"NVD", "CVSS", "CVEs"} {
			assert.StringNotContains(t, acronym, text)
		}
	}
}
