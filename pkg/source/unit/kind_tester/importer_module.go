package kind_tester

import "github.com/funtimecoding/soil/pkg/system/system_tester"

func (o *Tester) ImporterModule() string {
	o.t.Helper()
	directory := o.t.TempDir()
	system_tester.New(o.t).Write(
		directory,
		"go.mod",
		"module example\n\ngo 1.22\n",
	)
	system_tester.New(o.t).Write(
		directory,
		"model/model.go",
		"package model\n\nimport \"encoding/json\"\n\ntype Model struct{}\n\nfunc (Model) MarshalJSON() ([]byte, error) {\n\treturn json.Marshal(\"model\")\n}\n\nfunc (Model) MarshalText() ([]byte, error) {\n\treturn []byte(\"model\"), nil\n}\n\nfunc (Model) Error() string {\n\treturn \"model\"\n}\n\nfunc (Model) Write(text string) int {\n\treturn len(text)\n}\n",
	)

	return directory
}
