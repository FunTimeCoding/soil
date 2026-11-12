package module_tester

import "testing"

func RenamedIdentity(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
		File(
			"alfa/send_msg.go",
			"package alfa\n\nfunc SendMsg() string {\n\treturn MsgKey\n}\n",
		).
		File(
			"alfa/keys.go",
			"package alfa\n\nconst (\n\tMsgKey   = \"m\"\n\tOtherKey = \"o\"\n)\n",
		).
		File(
			"bravo/relay.go",
			"package bravo\n\nimport \"testmodule/alfa\"\n\nfunc Relay() string {\n\treturn alfa.SendMsg()\n}\n",
		).
		Directory()
}
