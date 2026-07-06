package scan

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"testing"
)

func TestCheckIdentityValid(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/constant/constant.go",
		"package constant\n\nimport \"github.com/funtimecoding/soil/pkg/identity\"\n\nvar Identity = identity.New(\"gotestd\", \"test\", \"gotestd\")\n",
	)
	assert.Integer(
		t,
		0,
		len(identityConcerns(v, "pkg/tool/gotestd", "gotestd")),
	)
}

func TestCheckIdentityMismatch(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/constant/constant.go",
		"package constant\n\nimport \"github.com/funtimecoding/soil/pkg/identity\"\n\nvar Identity = identity.New(\"wrong\", \"test\", \"wrong\")\n",
	)
	r := identityConcerns(v, "pkg/tool/gotestd", "gotestd")
	assert.Integer(t, 1, len(r))
	assert.String(t, IdentityMismatchKey, r[0].Key)
}

func TestCheckIdentityMissingFile(t *testing.T) {
	r := identityConcerns(
		virtual_file_system.New(),
		"pkg/tool/gotestd",
		"gotestd",
	)
	assert.Integer(t, 1, len(r))
	assert.String(t, IdentityMissingFileKey, r[0].Key)
}

func TestCheckIdentityMissingVariable(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/constant/constant.go",
		"package constant\n\nvar Name = \"test\"\n",
	)
	r := identityConcerns(v, "pkg/tool/gotestd", "gotestd")
	assert.Integer(t, 1, len(r))
	assert.String(t, IdentityMissingVariableKey, r[0].Key)
}
