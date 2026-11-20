package toolchain

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/fix_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenameMovesFileNamedForIdentity(t *testing.T) {
	m := module_tester.New(t)
	directory := m.RenamedIdentity()
	r := output.NewResultsWithDirectory(directory)
	m.RunDefault(m.Options(directory, false, false, "./..."), r)
	fix_tester.New(t).Result(
		fix_tester.Applied(r.Entries),
		"alfa/send_msg.go",
		"moved send_msg.go → send_message.go",
	)
	assert.False(
		t,
		system.FileExists(filepath.Join(directory, "alfa", "send_msg.go")),
	)
	assert.String(
		t,
		"package alfa\n\nfunc SendMessage() string {\n\treturn MessageKey\n}\n",
		testutil.ReadFile(
			t,
			filepath.Join(directory, "alfa", "send_message.go"),
		),
	)
	fix_tester.New(t).Builds(directory)
}

func TestRenameRealignsDeclarationBlock(t *testing.T) {
	m := module_tester.New(t)
	directory := m.RenamedIdentity()
	m.RunDefault(
		m.Options(directory, false, false, "./..."),
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package alfa\n\nconst (\n\tMessageKey = \"m\"\n\tOtherKey   = \"o\"\n)\n",
		testutil.ReadFile(t, filepath.Join(directory, "alfa", "keys.go")),
	)
}

func TestRenameKeepsFileWhoseNameIsTaken(t *testing.T) {
	m := module_tester.New(t)
	directory := m.RenamedIdentityTaken()
	r := output.NewResultsWithDirectory(directory)
	m.RunDefault(m.Options(directory, false, false, "./..."), r)
	fix_tester.New(t).Result(
		fix_tester.Blocked(r.Entries),
		"alfa/send_msg.go",
		"send_message.go is taken; send_msg.go keeps its name",
	)
	assert.String(
		t,
		"package alfa\n\nfunc SendMessage() string {\n\treturn \"x\"\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "alfa", "send_msg.go")),
	)
	assert.String(
		t,
		"package alfa\n\nfunc Deliver() string {\n\treturn \"y\"\n}\n",
		testutil.ReadFile(
			t,
			filepath.Join(directory, "alfa", "send_message.go"),
		),
	)
}

func TestRenameDiffShowsMoveWithoutWriting(t *testing.T) {
	m := module_tester.New(t)
	directory := m.RenamedIdentity()
	before := fix_tester.New(t).GoFiles(filepath.Join(directory, "alfa"))
	printed := assert.Capture(
		t,
		func() {
			m.RunDefault(
				m.Options(directory, true, false, "./..."),
				output.NewResultsWithDirectory(directory),
			)
		},
	)
	assert.Any(
		t,
		before,
		fix_tester.New(t).GoFiles(filepath.Join(directory, "alfa")),
	)
	assert.True(
		t,
		strings.Contains(
			printed,
			fmt.Sprintf(
				"--- %s\n+++ %s\n",
				filepath.Join(directory, "alfa", "send_msg.go"),
				filepath.Join(directory, "alfa", "send_message.go"),
			),
		),
	)
}
