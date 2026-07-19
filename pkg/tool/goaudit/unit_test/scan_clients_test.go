package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestClientsDiscovered(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/gitlab/client.go", "package gitlab\n")
	v.WriteString("pkg/gitlab/constant/constant.go", "package constant\n")
	c := scan.Clients(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(c))
	assert.String(t, "pkg/gitlab", c[0].Path)
	assert.Boolean(t, true, c[0].Constant)
}

func TestClientsExcludesTool(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotestd/client.go", "package gotestd\n")
	assert.Integer(
		t,
		0,
		len(scan.Clients(v, "test", scan.NewConfiguration())),
	)
}

func TestClientsExcludesConfiguration(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/relational/client.go", "package relational\n")
	assert.Integer(
		t,
		0,
		len(
			scan.Clients(
				v,
				"test",
				&scan.Configuration{Exclude: []string{"pkg/relational"}},
			),
		),
	)
}

func TestClientsNested(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/chat/mattermost/client.go", "package mattermost\n")
	c := scan.Clients(v, "test", scan.NewConfiguration())
	assert.Integer(t, 1, len(c))
	assert.String(t, "pkg/chat/mattermost", c[0].Path)
}
