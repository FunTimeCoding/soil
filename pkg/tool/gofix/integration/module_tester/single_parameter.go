package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester/module"

func (o *Tester) SingleParameter() string {
	o.t.Helper()

	return module.New(o.t,"example").
		File(
			"method.go",
			"package example\n\nimport \"context\"\n\ntype Client struct{}\n\nfunc (c *Client) Snapshot(\n\tx context.Context,\n) (string, error) {\n\treturn \"\", nil\n}\n",
		).
		File(
			"function.go",
			"package example\n\nfunc Process(\n\tname string,\n) error {\n\treturn nil\n}\n",
		).
		File(
			"two_params.go",
			"package example\n\nfunc TwoParams(\n\ta string,\n\tb string,\n) error {\n\treturn nil\n}\n",
		).
		File(
			"already_single.go",
			"package example\n\nfunc Short(x int) {}\n",
		).
		File(
			"too_long.go",
			"package example\n\nfunc VeryLongFunctionNameThatWouldExceedTheLimit(\n\tparameterWithAVeryLongName string,\n) error {\n\treturn nil\n}\n",
		).
		File(
			"expanded_function_type.go",
			"package example\n\nfunc Expanded(\n\tf func(\n\t\tint,\n\t\t[]byte,\n\t) error,\n) error {\n\treturn nil\n}\n",
		).
		File(
			"half_expanded_function_type.go",
			"package example\n\nfunc HalfExpanded(f func(\n\tint,\n\t[]byte,\n) error) error {\n\treturn nil\n}\n",
		).
		File(
			"commented_function_type.go",
			"package example\n\nfunc Commented(\n\tf func(\n\t\tint, // status\n\t\t[]byte,\n\t) error,\n) error {\n\treturn nil\n}\n",
		).
		File(
			"struct_type.go",
			"package example\n\nfunc Structured(\n\tv struct {\n\t\tA int\n\t\tB int\n\t},\n) error {\n\treturn nil\n}\n",
		).
		Directory()
}
