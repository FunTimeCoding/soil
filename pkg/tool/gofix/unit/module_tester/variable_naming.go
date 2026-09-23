package module_tester

import "testing"

func VariableNaming(t *testing.T) string {
	t.Helper()

	return New(t, "example").
		File(
			"wrong_single.go",
			"package example\n\nimport \"fmt\"\n\nfunc WrongSingle() {\n\tx := fmt.Errorf(\"test\")\n\t_ = x\n}\n",
		).
		File(
			"error_renamed.go",
			"package example\n\nimport \"fmt\"\n\nfunc ErrorRenamed() {\n\terr := fmt.Errorf(\"test\")\n\t_ = err\n}\n",
		).
		File(
			"error_chain_renamed.go",
			"package example\n\nimport \"fmt\"\n\nfunc ErrorChainRenamed() {\n\terr := fmt.Errorf(\"first\")\n\terr2 := fmt.Errorf(\"second\")\n\t_ = err\n\t_ = err2\n}\n",
		).
		File(
			"correct.go",
			"package example\n\nimport \"fmt\"\n\nfunc CorrectUntouched() {\n\te := fmt.Errorf(\"test\")\n\ts := \"hello\"\n\t_ = e\n\t_ = s\n}\n",
		).
		Directory()
}
