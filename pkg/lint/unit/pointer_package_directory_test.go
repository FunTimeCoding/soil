package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestPackageDirectory(t *testing.T) {
	assert.String(
		t,
		"pkg/provision/salt",
		pointer.PackageDirectory("pkg/provision/salt.Client"),
	)
	assert.String(
		t,
		"pkg/system/run",
		pointer.PackageDirectory("pkg/system/run/New()"),
	)
	assert.String(
		t,
		"pkg/face",
		pointer.PackageDirectory("pkg/face/EventNotifier"),
	)
	assert.String(
		t,
		"container/gopostgres",
		pointer.PackageDirectory("container/gopostgres/Containerfile"),
	)
	assert.String(
		t,
		"pkg/system/run",
		pointer.PackageDirectory("pkg/system/run.Runner.Start"),
	)
	assert.String(
		t,
		"doc/ai/runbook",
		pointer.PackageDirectory("${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/Lint"),
	)
	assert.String(
		t,
		"pkg/lint/pointer",
		pointer.PackageDirectory("pkg/lint/pointer/Names.ResolvePackage"),
	)
	assert.String(
		t,
		"pkg/web",
		pointer.PackageDirectory("pkg/web/RecoveryMiddleware"),
	)
	assert.String(
		t,
		"../soil/pkg/provision/salt",
		pointer.PackageDirectory("../soil/pkg/provision/salt.Client"),
	)
	assert.String(t, "", pointer.PackageDirectory("Symbol"))
	assert.String(t, "fmt", pointer.PackageDirectory("fmt.Println"))
	assert.String(t, "strings", pointer.PackageDirectory("strings"))
	assert.String(t, "crypto/x509", pointer.PackageDirectory("crypto/x509"))
	assert.String(
		t,
		"pkg/provision/salt",
		pointer.PackageDirectory("pkg/provision/salt"),
	)
}
