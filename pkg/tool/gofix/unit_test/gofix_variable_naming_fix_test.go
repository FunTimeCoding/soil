package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"path/filepath"
	"testing"
)

func TestVariableNamingFix(t *testing.T) {
	directory := writeVariableNamingTestModule(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunVariableNamingFixWithDirectory(
		[]string{"./..."},
		directory,
		false,
		r,
	)
	t.Run(
		"WrongSingleLetter",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"fmt\"\n\nfunc WrongSingle() {\n\te := fmt.Errorf(\"test\")\n\t_ = e\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "wrong_single.go"),
				),
			)
		},
	)
	t.Run(
		"ErrorRenamed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"fmt\"\n\nfunc ErrorRenamed() {\n\te := fmt.Errorf(\"test\")\n\t_ = e\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "error_renamed.go"),
				),
			)
		},
	)
	t.Run(
		"ErrorChainRenamed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"fmt\"\n\nfunc ErrorChainRenamed() {\n\te := fmt.Errorf(\"first\")\n\tf := fmt.Errorf(\"second\")\n\t_ = e\n\t_ = f\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(
						directory,
						"error_chain_renamed.go",
					),
				),
			)
		},
	)
	t.Run(
		"CorrectUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"fmt\"\n\nfunc CorrectUntouched() {\n\te := fmt.Errorf(\"test\")\n\ts := \"hello\"\n\t_ = e\n\t_ = s\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "correct.go"),
				),
			)
		},
	)
	t.Run(
		"ResultEntries",
		func(t *testing.T) {
			applied := filterApplied(r.Entries)
			assertResult(t, applied, "wrong_single.go", "renamed x → e")
			assertResult(t, applied, "error_renamed.go", "renamed err → e")
			assertResult(
				t,
				applied,
				"error_chain_renamed.go",
				"renamed err → e",
			)
			assertResult(
				t,
				applied,
				"error_chain_renamed.go",
				"renamed err2 → f",
			)
		},
	)
}

func writeVariableNamingTestModule(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteFile(
		t,
		directory,
		"go.mod",
		"module example\n\ngo 1.22\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"wrong_single.go",
		"package example\n\nimport \"fmt\"\n\nfunc WrongSingle() {\n\tx := fmt.Errorf(\"test\")\n\t_ = x\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"error_renamed.go",
		"package example\n\nimport \"fmt\"\n\nfunc ErrorRenamed() {\n\terr := fmt.Errorf(\"test\")\n\t_ = err\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"error_chain_renamed.go",
		"package example\n\nimport \"fmt\"\n\nfunc ErrorChainRenamed() {\n\terr := fmt.Errorf(\"first\")\n\terr2 := fmt.Errorf(\"second\")\n\t_ = err\n\t_ = err2\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"correct.go",
		"package example\n\nimport \"fmt\"\n\nfunc CorrectUntouched() {\n\te := fmt.Errorf(\"test\")\n\ts := \"hello\"\n\t_ = e\n\t_ = s\n}\n",
	)

	return directory
}
