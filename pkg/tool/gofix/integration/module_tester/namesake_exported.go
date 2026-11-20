package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) NamesakeExported() string {
	o.t.Helper()

	return module.New(o.t,"testmodule").
		File(
			"alfa/send.go",
			"package alfa\n\nfunc SendMsg() string {\n\treturn \"x\"\n}\n",
		).
		File(
			"bravo/relay.go",
			"package bravo\n\nimport \"testmodule/alfa\"\n\nfunc Relay() string {\n\treturn alfa.SendMsg()\n}\n",
		).
		File(
			"charlie/own.go",
			"package charlie\n\nfunc SendMsg() string {\n\treturn \"c\"\n}\n",
		).
		File(
			"delta/other.go",
			"package delta\n\nfunc Other() string {\n\treturn \"d\"\n}\n",
		).
		File(
			"echo/other.go",
			"package echo\n\nfunc Other() string {\n\treturn \"e\"\n}\n",
		).
		Directory()
}
