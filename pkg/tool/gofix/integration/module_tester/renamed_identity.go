package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) RenamedIdentity() string {
	o.t.Helper()

	return module.New(o.t,"testmodule").
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
