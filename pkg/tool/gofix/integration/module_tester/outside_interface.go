package module_tester

import "testing"

func OutsideInterface(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
		File(
			"alfa/sender.go",
			"package alfa\n\ntype Sender struct{}\n\nfunc (Sender) SendMsg() string {\n\treturn \"x\"\n}\n",
		).
		File(
			"bravo/messenger.go",
			"package bravo\n\nimport \"testmodule/alfa\"\n\ntype messenger interface {\n\tSendMsg() string\n}\n\nvar Chosen messenger = alfa.Sender{}\n",
		).
		Directory()
}
