package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"net/url"
	"strings"
	"testing"
)

func TestFlagsPageFiltersByState(t *testing.T) {
	v := base.New(t).Served()
	v.PostBody(
		t,
		constant.FlagPath,
		url.Values{"identifier": {"CVE-2099-10000"}},
	)
	v.PostBody(
		t,
		constant.FlagPath,
		url.Values{"identifier": {"CVE-2099-11111"}},
	)
	v.PostBody(
		t,
		constant.MitigatePath,
		url.Values{"identifier": {"CVE-2099-11111"}},
	)
	open := v.Body(t, "/?show=open")
	assert.StringContains(t, "CVE-2099-10000", open)
	assert.False(t, strings.Contains(open, "CVE-2099-11111"))
	mitigated := v.Body(t, "/?show=mitigated")
	assert.StringContains(t, "CVE-2099-11111", mitigated)
	assert.False(t, strings.Contains(mitigated, "CVE-2099-10000"))
	all := v.Body(t, "/?show=all")
	assert.StringContains(t, "CVE-2099-10000", all)
	assert.StringContains(t, "CVE-2099-11111", all)
}

func TestFlagRowOpensAndCloses(t *testing.T) {
	v := base.New(t).Served()
	v.PostBody(
		t,
		constant.FlagPath,
		url.Values{
			"identifier":       {vulnerabilityConstant.KnownIdentifier},
			constant.NoteField: {"quebec"},
		},
	)
	query := url.Values{
		"identifier": {vulnerabilityConstant.KnownIdentifier},
	}.Encode()
	opened := v.Body(t, join.Empty(constant.FlagOpenPath, "?", query))
	assert.StringContains(t, "Note: ", opened)
	assert.StringContains(t, "quebec", opened)
	assert.StringContains(t, "Full lookup", opened)
	assert.StringContains(t, `action="/mitigate"`, opened)
	closed := v.Body(t, join.Empty(constant.FlagShutPath, "?", query))
	assert.False(t, strings.Contains(closed, "Full lookup"))
	assert.StringContains(t, "CVE-2024-3094", closed)
}
