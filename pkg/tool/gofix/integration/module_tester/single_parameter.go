package module_tester

import "testing"

func SingleParameter(t *testing.T) string {
	t.Helper()

	return New(t, "example").
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
		Directory()
}
