package module_tester

import "testing"

func ImportAliasSuggestion(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
		File(
			"pkg/time/time.go",
			"package time\n\nfunc Stamp() string { return \"stamp\" }\n",
		).
		File(
			"pkg/delta/clock.go",
			"package delta\n\nimport (\n\tlibrary \"testmodule/pkg/time\"\n\t\"time\"\n)\n\nfunc Clock() string {\n\treturn library.Stamp() + time.UTC.String()\n}\n",
		).
		File(
			"pkg/hotel/string/string.go",
			"package string\n\nfunc Hotel() int { return 1 }\n",
		).
		File(
			"pkg/india/string/string.go",
			"package string\n\nfunc India() int { return 2 }\n",
		).
		File(
			"pkg/hotel/builtin.go",
			"package hotel\n\nimport (\n\town \"testmodule/pkg/hotel/string\"\n\tindia \"testmodule/pkg/india/string\"\n)\n\nfunc Builtin() int {\n\treturn own.Hotel() + india.India()\n}\n",
		).
		Directory()
}
