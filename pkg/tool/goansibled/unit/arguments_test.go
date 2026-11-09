package unit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goansibled"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"testing"
)

func TestPlaybookComesFromEnvironment(t *testing.T) {
	t.Setenv(
		constant.AnsiblePlaybookEnvironment,
		"playbooks/a.yaml,playbooks/b.yaml",
	)
	a := argument.NewInstance(constant.Identity)
	goansibled.Arguments(a)
	assert.Nil(t, a.ParseArguments(nil))
	assert.Strings(
		t,
		[]string{"playbooks/a.yaml", "playbooks/b.yaml"},
		a.Slice(argumentConstant.Playbook),
	)
}

func TestPlaybookFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(constant.AnsiblePlaybookEnvironment, "playbooks/a.yaml")
	a := argument.NewInstance(constant.Identity)
	goansibled.Arguments(a)
	assert.Nil(t, a.ParseArguments([]string{"--playbook", "playbooks/c.yaml"}))
	assert.Strings(
		t,
		[]string{"playbooks/c.yaml"},
		a.Slice(argumentConstant.Playbook),
	)
}
