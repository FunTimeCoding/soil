package module_tester

import "testing"

func ReplacedLibrary(t *testing.T) string {
	t.Helper()

	return New(t, "other.test/lib").
		File(
			"alfa/send.go",
			"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		).
		File(
			"bravo/relay.go",
			"package bravo\n\nimport \"other.test/lib/alfa\"\n\nfunc Relay() string {\n\treturn alfa.SendMsg()\n}\n",
		).
		Directory()
}
