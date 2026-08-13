package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"path/filepath"
	"testing"
)

func TestCompositeFormatFix(t *testing.T) {
	directory := writeCompositeFormatTestModule(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunCompositeFormatFixWithDirectory([]string{"./..."}, directory, r)
	t.Run(
		"SharedLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc SharedLine() []int {\n\treturn []int{1, 2, 3}\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "shared_line.go")),
			)
		},
	)
	t.Run(
		"FirstOnBraceLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc FirstOnBraceLine() []string {\n\treturn []string{\"alpha\", \"bravo\"}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "first_on_brace_line.go"),
				),
			)
		},
	)
	t.Run(
		"NestedCalls",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Item struct {\n\tValue int\n}\n\nfunc NewItem(v int) *Item {\n\treturn &Item{Value: v}\n}\n\nfunc NestedCalls() []*Item {\n\treturn []*Item{\n\t\tNewItem(1),\n\t\tNewItem(2),\n\t}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "nested_calls.go"),
				),
			)
		},
	)
	t.Run(
		"CompliantCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Compliant() []int {\n\treturn []int{1, 2, 3}\n}\n\nfunc CompliantMultiLine() []int {\n\treturn []int{1, 2, 3}\n}\n",
				testutil.ReadFile(t, filepath.Join(directory, "compliant.go")),
			)
		},
	)
	t.Run(
		"CollapseShortSlice",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc CollapseShortSlice() []int {\n\treturn []int{1, 2, 3}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "collapse_short_slice.go"),
				),
			)
		},
	)
	t.Run(
		"SingleLineLongSplits",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc SingleLineLong() {\n\tif true {\n\t\tif true {\n\t\t\t_ = []string{\n\t\t\t\t\"aaa\",\n\t\t\t\t\"bbb\",\n\t\t\t\t\"ccc\",\n\t\t\t\t\"ddd\",\n\t\t\t\t\"eee\",\n\t\t\t\t\"fff\",\n\t\t\t\t\"ggg\",\n\t\t\t\t\"hhh\",\n\t\t\t\t\"iii\",\n\t\t\t}\n\t\t}\n\t}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "single_line_long.go"),
				),
			)
		},
	)
	t.Run(
		"MapCollapses",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc MapCollapses() map[string]int {\n\treturn map[string]int{\"x\": 1, \"y\": 2}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "map_collapses.go"),
				),
			)
		},
	)
	t.Run(
		"MultiLineElementCollapsed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc MultiLineElement() [][]int {\n\treturn [][]int{\n\t\t{1, 2},\n\t\t{3},\n\t}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "multi_line_element.go"),
				),
			)
		},
	)
	t.Run(
		"CollapsePreservesLong",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc CollapsePreservesLong() []string {\n\treturn []string{\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "collapse_preserves_long.go"),
				),
			)
		},
	)
	t.Run(
		"ResultEntries",
		func(t *testing.T) {
			applied := filterApplied(r.Entries)
			assert.Integer(t, 9, len(applied))
			assertResult(
				t,
				applied,
				"shared_line.go",
				"formatted composite literal (line 4)",
			)
			assertResult(
				t,
				applied,
				"first_on_brace_line.go",
				"formatted composite literal (line 4)",
			)
			assertResult(
				t,
				applied,
				"nested_calls.go",
				"formatted composite literal (line 12)",
			)
			assertResult(
				t,
				applied,
				"collapse_short_slice.go",
				"formatted composite literal (line 4)",
			)
			assertResult(
				t,
				applied,
				"compliant.go",
				"formatted composite literal (line 8)",
			)
			assertResult(
				t,
				applied,
				"single_line_long.go",
				"formatted composite literal (line 6)",
			)
			assertResult(
				t,
				applied,
				"map_collapses.go",
				"formatted composite literal (line 4)",
			)
			assertResult(
				t,
				applied,
				"multi_line_element.go",
				"formatted composite literal (line 5)",
			)
		},
	)
}

func writeCompositeFormatTestModule(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteFile(t, directory, "go.mod", "module example\n\ngo 1.22\n")
	testutil.WriteFile(
		t,
		directory,
		"shared_line.go",
		"package example\n\nfunc SharedLine() []int {\n\treturn []int{\n\t\t1, 2,\n\t\t3,\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"first_on_brace_line.go",
		"package example\n\nfunc FirstOnBraceLine() []string {\n\treturn []string{\"alpha\",\n\t\t\"bravo\",\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"nested_calls.go",
		"package example\n\ntype Item struct {\n\tValue int\n}\n\nfunc NewItem(v int) *Item {\n\treturn &Item{Value: v}\n}\n\nfunc NestedCalls() []*Item {\n\treturn []*Item{\n\t\tNewItem(1), NewItem(\n\t\t\t2,\n\t\t),\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"compliant.go",
		"package example\n\nfunc Compliant() []int {\n\treturn []int{1, 2, 3}\n}\n\nfunc CompliantMultiLine() []int {\n\treturn []int{\n\t\t1,\n\t\t2,\n\t\t3,\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"single_line_long.go",
		"package example\n\nfunc SingleLineLong() {\n\tif true {\n\t\tif true {\n\t\t\t_ = []string{\"aaa\", \"bbb\", \"ccc\", \"ddd\", \"eee\", \"fff\", \"ggg\", \"hhh\", \"iii\"}\n\t\t}\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"map_collapses.go",
		"package example\n\nfunc MapCollapses() map[string]int {\n\treturn map[string]int{\n\t\t\"x\": 1,\n\t\t\"y\": 2,\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"multi_line_element.go",
		"package example\n\nfunc MultiLineElement() [][]int {\n\treturn [][]int{\n\t\t{\n\t\t\t1,\n\t\t\t2,\n\t\t},\n\t\t{3},\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"collapse_short_slice.go",
		"package example\n\nfunc CollapseShortSlice() []int {\n\treturn []int{\n\t\t1,\n\t\t2,\n\t\t3,\n\t}\n}\n",
	)
	testutil.WriteFile(
		t,
		directory,
		"collapse_preserves_long.go",
		"package example\n\nfunc CollapsePreservesLong() []string {\n\treturn []string{\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t}\n}\n",
	)

	return directory
}
