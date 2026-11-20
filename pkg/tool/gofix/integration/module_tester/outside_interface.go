package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) OutsideInterface() string {
	o.t.Helper()

	return module.New(o.t,"testmodule").
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
