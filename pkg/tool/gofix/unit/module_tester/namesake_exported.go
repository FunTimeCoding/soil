package module_tester

import "testing"

func NamesakeExported(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
		File(
			"alfa/send.go",
			"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		).
		File(
			"bravo/relay.go",
			"package bravo\n\nimport \"testmodule/alfa\"\n\nfunc Relay() string {\n\treturn alfa.SendMsg()\n}\n",
		).
		File(
			"charlie/own.go",
			"package charlie\n\nfunc SendMsg() string {\n\treturn \"c\"\n}\n",
		).
		File(
			"delta/other.go",
			"package delta\n\nfunc Other() string {\n\treturn \"d\"\n}\n",
		).
		File(
			"echo/other.go",
			"package echo\n\nfunc Other() string {\n\treturn \"e\"\n}\n",
		).
		Directory()
}
