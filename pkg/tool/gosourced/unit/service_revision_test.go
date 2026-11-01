package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/revision"
	"testing"
)

func TestRevisionQualifiesImportedAndOwnReferences(t *testing.T) {
	r := revision.New(
		map[string][]byte{
			"go.mod": []byte("module example.org/alfa\n"),
			"bravo/bravo.go": []byte(
				"package bravo\n\nimport (\n\t\"example.org/alfa/charlie\"\n\t\"example.org/delta/go-echo/v2\"\n)\n\nconst Limit = 3\n\nfunc Run(count int) int {\n\ttotal := charlie.Sum(count, Limit)\n\treturn echo.Clamp(total)\n}\n",
			),
			"charlie/charlie.go": []byte(
				"package charlie\n\nfunc Sum(a int, b int) int {\n\treturn a + b\n}\n",
			),
		},
	)
	run := revisionSymbol(t, r,"bravo bravo Run")
	assert.String(
		t,
		"func Run(count int) int {\n\ttotal := \"example.org/alfa/charlie\".Sum(count, \"example.org/alfa/bravo\".Limit)\n\treturn \"example.org/delta/go-echo/v2\".Clamp(total)\n}",
		run.Body,
	)
	assert.String(t, "func Run(count int) int", run.Signature)
	assert.String(t, "bravo/bravo.go", run.File)
	assert.Integer(t, 10, run.Line)
}

func TestRevisionKeepsLocalsUnqualified(t *testing.T) {
	r := revision.New(
		map[string][]byte{
			"go.mod": []byte("module example.org/alfa\n"),
			"alfa.go": []byte(
				"package alfa\n\nvar count = 1\n\nfunc Next(count int) int {\n\treturn count + 1\n}\n\nfunc Last() int {\n\treturn count\n}\n",
			),
		},
	)
	assert.String(
		t,
		"func Next(count int) int {\n\treturn count + 1\n}",
		revisionSymbol(t, r,". alfa Next").Body,
	)
	assert.String(
		t,
		"func Last() int {\n\treturn \"example.org/alfa\".count\n}",
		revisionSymbol(t, r,". alfa Last").Body,
	)
}

func TestRevisionKeysMethodsByReceiver(t *testing.T) {
	r := revision.New(
		map[string][]byte{
			"go.mod": []byte("module example.org/alfa\n"),
			"alfa.go": []byte(
				"package alfa\n\ntype Tank struct {\n\tLevel int\n}\n\nfunc (t *Tank) Fill(amount int) {\n\tt.Level += amount\n}\n",
			),
		},
	)
	assert.Strings(
		t,
		[]string{". alfa Tank", ". alfa Tank.Fill"},
		revisionKeys(r),
	)
	fill := revisionSymbol(t, r,". alfa Tank.Fill")
	assert.String(t, "method", fill.Kind)
	assert.String(
		t,
		"func (t *\"example.org/alfa\".Tank) Fill(amount int)",
		fill.Signature,
	)
}

func TestRevisionCountsIotaPositions(t *testing.T) {
	r := revision.New(
		map[string][]byte{
			"go.mod": []byte("module example.org/alfa\n"),
			"alfa.go": []byte(
				"package alfa\n\ntype Mode int\n\nconst (\n\tIdle Mode = iota\n\tBusy\n)\n",
			),
		},
	)
	assert.String(
		t,
		"const Busy \"example.org/alfa\".Mode = iota // iota 1",
		revisionSymbol(t, r,". alfa Busy").Body,
	)
}

func TestRevisionDistinguishesRepeatedNames(t *testing.T) {
	r := revision.New(
		map[string][]byte{
			"go.mod":  []byte("module example.org/alfa\n"),
			"alfa.go": []byte("package alfa\n\nfunc init() {}\n"),
			"bravo.go": []byte(
				"package alfa\n\nfunc init() {}\n\nfunc init() {}\n",
			),
		},
	)
	assert.Strings(
		t,
		[]string{
			". alfa init @alfa.go",
			". alfa init @bravo.go#1",
			". alfa init @bravo.go#2",
		},
		revisionKeys(r),
	)
}

func TestRevisionReportsUnparsedFiles(t *testing.T) {
	r := revision.New(
		map[string][]byte{
			"go.mod":    []byte("module example.org/alfa\n"),
			"alfa.go":   []byte("package alfa\n\nconst Bravo = 1\n"),
			"broken.go": []byte("package alfa\n\nfunc {\n"),
		},
	)
	assert.Strings(t, []string{"broken.go"}, r.Unparsed)
	assert.Strings(t, []string{". alfa Bravo"}, revisionKeys(r))
}
