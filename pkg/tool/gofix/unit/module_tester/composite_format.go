package module_tester

import "testing"

func CompositeFormat(t *testing.T) string {
	t.Helper()

	return New(t, "example").
		File(
			"shared_line.go",
			"package example\n\nfunc SharedLine() []int {\n\treturn []int{\n\t\t1, 2,\n\t\t3,\n\t}\n}\n",
		).
		File(
			"first_on_brace_line.go",
			"package example\n\nfunc FirstOnBraceLine() []string {\n\treturn []string{\"alpha\",\n\t\t\"bravo\",\n\t}\n}\n",
		).
		File(
			"nested_calls.go",
			"package example\n\ntype Item struct {\n\tValue int\n}\n\nfunc NewItem(v int) *Item {\n\treturn &Item{Value: v}\n}\n\nfunc NestedCalls() []*Item {\n\treturn []*Item{\n\t\tNewItem(1), NewItem(\n\t\t\t2,\n\t\t),\n\t}\n}\n",
		).
		File(
			"compliant.go",
			"package example\n\nfunc Compliant() []int {\n\treturn []int{1, 2, 3}\n}\n\nfunc CompliantMultiLine() []int {\n\treturn []int{\n\t\t1,\n\t\t2,\n\t\t3,\n\t}\n}\n",
		).
		File(
			"single_line_long.go",
			"package example\n\nfunc SingleLineLong() {\n\tif true {\n\t\tif true {\n\t\t\t_ = []string{\"aaa\", \"bbb\", \"ccc\", \"ddd\", \"eee\", \"fff\", \"ggg\", \"hhh\", \"iii\"}\n\t\t}\n\t}\n}\n",
		).
		File(
			"map_collapses.go",
			"package example\n\nfunc MapCollapses() map[string]int {\n\treturn map[string]int{\n\t\t\"x\": 1,\n\t\t\"y\": 2,\n\t}\n}\n",
		).
		File(
			"multi_line_element.go",
			"package example\n\nfunc MultiLineElement() [][]int {\n\treturn [][]int{\n\t\t{\n\t\t\t1,\n\t\t\t2,\n\t\t},\n\t\t{3},\n\t}\n}\n",
		).
		File(
			"collapse_short_slice.go",
			"package example\n\nfunc CollapseShortSlice() []int {\n\treturn []int{\n\t\t1,\n\t\t2,\n\t\t3,\n\t}\n}\n",
		).
		File(
			"collapse_preserves_long.go",
			"package example\n\nfunc CollapsePreservesLong() []string {\n\treturn []string{\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t}\n}\n",
		).
		Directory()
}
