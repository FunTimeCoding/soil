package module_tester

import "testing"

func UnexportedWithTests(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
		File(
			"alfa/send.go",
			"package alfa\n\nfunc sendMsg() string {\n\treturn \"x\"\n}\n",
		).
		File(
			"alfa/send_test.go",
			"package alfa\n\nimport \"testing\"\n\nfunc TestSend(t *testing.T) {\n\t_ = sendMsg()\n}\n",
		).
		Directory()
}
