package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"testing"
)

func TestMockListsSeededRemediationLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	x := t.Context()
	m.SeedRemediationItem(
		client.RemediationItem{Subject: "upgrade:trixie:victor", TopScore: 9.8},
	)
	m.SeedRemediationItem(
		client.RemediationItem{Subject: "CVE-2099-55001", TopScore: 7.5},
	)
	m.SeedRemediationItem(
		client.RemediationItem{Subject: "release:golf", Unknown: true},
	)
	m.SeedRemediationItem(
		client.RemediationItem{
			Subject:   "reboot:foxtrot",
			TopScore:  9.1,
			Addressed: true,
		},
	)
	critical, e := m.ListRemediation(x, nil)
	assert.FatalOnError(t, e)
	assert.Count(t, 3, critical.Items)
	assert.Integer(t, 2, critical.Open)
	assert.Integer(t, 1, critical.Addressed)
	high := 7.0
	all, f := m.ListRemediation(x, &client.ListRemediationParams{Score: &high})
	assert.FatalOnError(t, f)
	assert.Count(t, 4, all.Items)
}
