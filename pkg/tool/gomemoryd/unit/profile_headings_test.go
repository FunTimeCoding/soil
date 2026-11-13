package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/profile"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func TestProfileHeadingsReadBackServedTiers(t *testing.T) {
	o := service_tester.New(t)
	standing := createTagged(t, o, "standing rule", constant.AlwaysTag)
	retry, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	result, _, f := o.Service.Profile(constant.FixtureContent, "", true)
	assert.FatalOnError(t, f)
	always := profile.Headings(result.Text, constant.AlwaysSectionHeading)
	assert.Count(t, 1, always)
	assert.Integer(t, int(standing.Identifier), int(always[0].Identifier))
	assert.String(t, "standing rule", always[0].Name)
	relevant := profile.Headings(result.Text, constant.RelevantSectionHeading)
	assert.Count(t, 1, relevant)
	assert.Integer(t, int(retry.Identifier), int(relevant[0].Identifier))
	assert.String(t, "retry policy", relevant[0].Name)
}

func TestProfileHeadingsSkipIndexTier(t *testing.T) {
	o := service_tester.New(t)
	createTagged(t, o, "standing rule", constant.AlwaysTag)
	_, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	result, _, f := o.Service.Profile("", "", true)
	assert.FatalOnError(t, f)
	assert.StringContains(t, "retry policy", result.Text)
	assert.Count(
		t,
		0,
		profile.Headings(result.Text, constant.IndexSectionHeading),
	)
}

func TestProfileHeadingKeepsParenthesesAndDropsTags(t *testing.T) {
	m := record.NewMemory()
	m.Identifier = 12
	m.Name = "rule (draft)"
	m.Content = "## not a heading\n(5)"
	m.Tags = []string{"always", "inner"}
	result := profile.NewResult([]record.Memory{*m}, nil, nil, nil, nil)
	always := profile.Headings(result.Text, constant.AlwaysSectionHeading)
	assert.Count(t, 1, always)
	assert.Integer(t, 12, int(always[0].Identifier))
	assert.String(t, "rule (draft)", always[0].Name)
}
