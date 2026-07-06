package scan

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"testing"
)

func TestClientsDiscovered(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/gitlab/client.go", "package gitlab\n")
	v.WriteString("pkg/gitlab/constant/constant.go", "package constant\n")
	c := Clients(v, "test", &Configuration{})
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
		len(Clients(v, "test", &Configuration{})),
	)
}

func TestClientsExcludesConfiguration(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/bolt/client.go", "package bolt\n")
	assert.Integer(
		t,
		0,
		len(
			Clients(
				v,
				"test",
				&Configuration{Exclude: []string{"pkg/bolt"}},
			),
		),
	)
}

func TestClientsNested(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/chat/mattermost/client.go", "package mattermost\n")
	c := Clients(v, "test", &Configuration{})
	assert.Integer(t, 1, len(c))
	assert.String(t, "pkg/chat/mattermost", c[0].Path)
}
