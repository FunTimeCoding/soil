package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) UnexportedWithTests() string {
	o.t.Helper()

	return module.New(o.t,"testmodule").
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
