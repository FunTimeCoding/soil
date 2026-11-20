package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) RenamedIdentityTaken() string {
	o.t.Helper()

	return module.New(o.t,"testmodule").
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
