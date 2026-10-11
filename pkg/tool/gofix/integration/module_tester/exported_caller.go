package module_tester

import "testing"

func ExportedCaller(t *testing.T) string {
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
		Directory()
}
