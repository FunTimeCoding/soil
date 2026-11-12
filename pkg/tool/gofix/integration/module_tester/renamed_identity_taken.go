package module_tester

import "testing"

func RenamedIdentityTaken(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
		File(
			"alfa/send_msg.go",
			"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		).
		File(
			"alfa/send_message.go",
			"package alfa\n\nfunc Deliver() string {\n\treturn \"y\"\n}\n",
		).
		Directory()
}
