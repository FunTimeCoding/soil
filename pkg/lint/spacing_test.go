package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"strings"
	"testing"
)

func TestSpacingBlankBeforeControl(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.MissingBlankBeforeControlKey,
				Text:     constant.MissingBlankBeforeControlText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     5,
				LineText: "\tif x > 0 {",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeReturn(t *testing.T) {
	l := Spacing(
		upper.Bravo,
		strings.NewReader(
			"package example\n\nfunc Example() int {\n\tx := 1\n\treturn x\n}\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      constant.MissingBlankBeforeExitKey,
				Text:     constant.MissingBlankBeforeExitText,
				Path:     "Bravo",
				Type:     concern.Line,
				Line:     5,
				LineText: "\treturn x",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() int {\n\tx := 1\n\n\treturn x\n}\n",
		l,
	)
}

func TestSpacingBlankAfterControl(t *testing.T) {
	l := Spacing(
		upper.Charlie,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\tfmt.Println(\"b\")\n}\n",
		),
	)
	assertReport(
		t,
		"Charlie",
		true,
		[]*concern.Concern{
			{
				Key:      constant.MissingBlankAfterControlKey,
				Text:     constant.MissingBlankAfterControlText,
				Path:     "Charlie",
				Type:     concern.Line,
				Line:     9,
				LineText: "\tfmt.Println(\"b\")",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		l,
	)
}

func TestSpacingReturnFirstInBlock(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() int {\n\treturn 1\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() int {\n\treturn 1\n}\n",
		l,
	)
}

func TestSpacingControlFirstInBlock(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n}\n",
		l,
	)
}

func TestSpacingConsecutiveClosingBraces(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t}\n}\n",
		l,
	)
}

func TestSpacingElse(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t} else {\n\t\tfmt.Println(\"b\")\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t} else {\n\t\tfmt.Println(\"b\")\n\t}\n}\n",
		l,
	)
}

func TestSpacingExtraneousBlanks(t *testing.T) {
	l := Spacing(
		upper.Delta,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\n\ty := 2\n}\n",
		),
	)
	assertReport(
		t,
		"Delta",
		true,
		[]*concern.Concern{
			{
				Key:      constant.ExtraneousBlankLineKey,
				Text:     constant.ExtraneousBlankLineText,
				Path:     "Delta",
				Type:     concern.Line,
				Line:     6,
				LineText: "",
				Fixed:    true,
			},
			{
				Key:      constant.BlankInsideFunctionKey,
				Text:     constant.BlankInsideFunctionText,
				Path:     "Delta",
				Type:     concern.Line,
				Line:     5,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\ty := 2\n}\n",
		l,
	)
}

func TestSpacingClean(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tfmt.Println(\"a\")\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tfmt.Println(\"a\")\n}\n",
		l,
	)
}

func TestSpacingClosingBraceBeforeDefault(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example(x int) {\n\tswitch x {\n\tcase 1:\n\t\tif x > 0 {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\tdefault:\n\t\tfmt.Println(\"b\")\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example(x int) {\n\tswitch x {\n\tcase 1:\n\t\tif x > 0 {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\tdefault:\n\t\tfmt.Println(\"b\")\n\t}\n}\n",
		l,
	)
}

func TestSpacingClosingBraceBeforeComma(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example(run func()) {\n\trun()\n}\n\nfunc Call() {\n\tExample(func() {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t})\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example(run func()) {\n\trun()\n}\n\nfunc Call() {\n\tExample(func() {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t})\n}\n",
		l,
	)
}

func TestSpacingContinueAsIdentifierPrefix(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example(\n\tcontinueOnError bool,\n) {\n\tfmt.Println(continueOnError)\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example(\n\tcontinueOnError bool,\n) {\n\tfmt.Println(continueOnError)\n}\n",
		l,
	)
}

func TestSpacingCommentBeforeControl(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\t// Check condition\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := 1\n\t// Check condition\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		l,
	)
}

func TestSpacingCompositeLiteralInFunction(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := map[string]int{\n\t\t\"a\": 1,\n\t}\n\ty := 2\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := map[string]int{\n\t\t\"a\": 1,\n\t}\n\ty := 2\n}\n",
		l,
	)
}

func TestSpacingRawStringLiteral(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nvar Template = `\nif (x) {\n\treturn false;\n}\ny = 1;\n`\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nvar Template = `\nif (x) {\n\treturn false;\n}\ny = 1;\n`\n",
		l,
	)
}

func TestSpacingVariableBlockCompositeLiteral(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nvar (\n\tStatuses = []string{\n\t\t\"open\",\n\t\t\"closed\",\n\t}\n\tOther = \"other\"\n)\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nvar (\n\tStatuses = []string{\n\t\t\"open\",\n\t\t\"closed\",\n\t}\n\tOther = \"other\"\n)\n",
		l,
	)
}

func TestSpacingBlankBetweenAssignments(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\ty := 2\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.BlankInsideFunctionKey,
				Text:     constant.BlankInsideFunctionText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     5,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\ty := 2\n}\n",
		l,
	)
}

func TestSpacingBlankAtStartOfFunction(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\n\tx := 1\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.BlankInsideFunctionKey,
				Text:     constant.BlankInsideFunctionText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     4,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n}\n",
		l,
	)
}

func TestSpacingBlankAtEndOfFunction(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.BlankInsideFunctionKey,
				Text:     constant.BlankInsideFunctionText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     5,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeReturnFirstInBlock(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() int {\n\n\treturn 1\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.BlankInsideFunctionKey,
				Text:     constant.BlankInsideFunctionText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     4,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() int {\n\treturn 1\n}\n",
		l,
	)
}

func TestSpacingBlankBetweenAssignmentsInNestedBlock(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\n\t\ty := 2\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.BlankInsideFunctionKey,
				Text:     constant.BlankInsideFunctionText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     6,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\t\ty := 2\n\t}\n}\n",
		l,
	)
}

func TestSpacingBlankAfterControlBeforeStatement(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeIfInFunction(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := 1\n\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		l,
	)
}

func TestSpacingBlankInFunctionLiteral(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tf := func() {\n\t\tx := 1\n\n\t\ty := 2\n\t}\n\tf()\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.BlankInsideFunctionKey,
				Text:     constant.BlankInsideFunctionText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     6,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tf := func() {\n\t\tx := 1\n\t\ty := 2\n\t}\n\tf()\n}\n",
		l,
	)
}

func TestSpacingCompositeLiteralCloserDoesNotLeakDepth(t *testing.T) {
	l := Spacing(
		"composite",
		strings.NewReader(
			"package example\n\nfunc First() {\n\t_ = []string{\n\t\t\"a\",\n\t}\n}\n\nfunc Second() {}\n",
		),
	)
	assertReport(
		t,
		"composite",
		false,
		nil,
		"package example\n\nfunc First() {\n\t_ = []string{\n\t\t\"a\",\n\t}\n}\n\nfunc Second() {}\n",
		l,
	)
}

func TestSpacingBlankBeforeCommentBeforeControl(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\t// check\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := 1\n\n\t// check\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeCommentBeforeAssignment(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\t// note\n\ty := 2\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := 1\n\n\t// note\n\ty := 2\n}\n",
		l,
	)
}

func TestSpacingMultiLineConditionBlankAfterControl(t *testing.T) {
	l := Spacing(
		"multiline",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif strings.Contains(\n\t\ts,\n\t\t\"x\",\n\t) {\n\t\tcontinue\n\t}\n\n\tx := 1\n\t_ = x\n}\n",
		),
	)
	assertReport(
		t,
		"multiline",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif strings.Contains(\n\t\ts,\n\t\t\"x\",\n\t) {\n\t\tcontinue\n\t}\n\n\tx := 1\n\t_ = x\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeDefer(t *testing.T) {
	l := Spacing(
		"defer",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx, e := f()\n\n\tdefer g()\n\t_ = x\n\t_ = e\n}\n",
		),
	)
	assertReport(
		t,
		"defer",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx, e := f()\n\n\tdefer g()\n\t_ = x\n\t_ = e\n}\n",
		l,
	)
}

func TestSpacingMissingBlankBetweenFunctions(t *testing.T) {
	l := Spacing(
		"decl",
		strings.NewReader(
			"package example\n\nfunc First() {}\nfunc Second() {}\n",
		),
	)
	assertReport(
		t,
		"decl",
		true,
		[]*concern.Concern{
			{
				Key:      constant.MissingBlankBeforeDeclarationKey,
				Text:     constant.MissingBlankBeforeDeclarationText,
				Path:     "decl",
				Type:     concern.Line,
				Line:     4,
				LineText: "func Second() {}",
				Fixed:    true,
			},
		},
		"package example\n\nfunc First() {}\n\nfunc Second() {}\n",
		l,
	)
}

func TestSpacingConsecutiveTopLevelVariableNoBlankAdded(t *testing.T) {
	l := Spacing(
		"topvar",
		strings.NewReader(
			"package example\n\nvar A = 1\nvar B = 2\nvar C = 3\n",
		),
	)
	assertReport(
		t,
		"topvar",
		false,
		nil,
		"package example\n\nvar A = 1\nvar B = 2\nvar C = 3\n",
		l,
	)
}

func TestSpacingSpuriousBlankBetweenTopLevelVarsRemoved(t *testing.T) {
	l := Spacing(
		"topvar2",
		strings.NewReader(
			"package example\n\nvar A = 1\n\nvar B = 2\n\nvar C = 3\n",
		),
	)
	assertReport(
		t,
		"topvar2",
		true,
		[]*concern.Concern{
			{
				Key:      constant.ExtraneousTopLevelBlankKey,
				Text:     constant.ExtraneousTopLevelBlankText,
				Path:     "topvar2",
				Type:     concern.Line,
				Line:     4,
				LineText: "",
				Fixed:    true,
			},
			{
				Key:      constant.ExtraneousTopLevelBlankKey,
				Text:     constant.ExtraneousTopLevelBlankText,
				Path:     "topvar2",
				Type:     concern.Line,
				Line:     6,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nvar A = 1\nvar B = 2\nvar C = 3\n",
		l,
	)
}

func TestSpacingBlankBeforeCommentAfterControlIdempotent(t *testing.T) {
	input := "package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\t\t_ = x\n\t}\n\n\t// note\n\ty := 2\n\t_ = y\n}\n"
	l := Spacing("idempotent", strings.NewReader(input))
	assertReport(t, "idempotent", false, nil, input, l)
}

func TestSpacingBlankAfterControlBeforeRegularStatementPreserved(t *testing.T) {
	l := Spacing(
		"ctrlblank",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\t\t_ = x\n\t}\n\n\ty := 2\n\t_ = y\n}\n",
		),
	)
	assertReport(
		t,
		"ctrlblank",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\t\t_ = x\n\t}\n\n\ty := 2\n\t_ = y\n}\n",
		l,
	)
}

func TestSpacingBlankAfterNestedControlBeforeRegularStatementPreserved(t *testing.T) {
	l := Spacing(
		"nestedctrl",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tfor _, s := range []string{\n\t\t\"a\",\n\t\t\"b\",\n\t} {\n\t\tif !contains(s) {\n\t\t\tadd(s)\n\t\t}\n\t}\n\n\tif !contains(\"c\") {\n\t\tadd(\"c\")\n\t}\n\n\tresult := count()\n\t_ = result\n}\n",
		),
	)
	assertReport(
		t,
		"nestedctrl",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tfor _, s := range []string{\n\t\t\"a\",\n\t\t\"b\",\n\t} {\n\t\tif !contains(s) {\n\t\t\tadd(s)\n\t\t}\n\t}\n\n\tif !contains(\"c\") {\n\t\tadd(\"c\")\n\t}\n\n\tresult := count()\n\t_ = result\n}\n",
		l,
	)
}

func TestSpacingBlankAfterCommentRemoved(t *testing.T) {
	l := Spacing(
		"aftercomment",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\t// note\n\n\tif true {\n\t\treturn\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"aftercomment",
		true,
		[]*concern.Concern{
			{
				Key:      constant.BlankInsideFunctionKey,
				Text:     constant.BlankInsideFunctionText,
				Path:     "aftercomment",
				Type:     concern.Line,
				Line:     5,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\t// note\n\tif true {\n\t\treturn\n\t}\n}\n",
		l,
	)
}

func TestSpacingInterFunctionBlankAfterNestedFuncLiteral(t *testing.T) {
	l := Spacing(
		"interfunc",
		strings.NewReader(
			"package example\n\nfunc First() {\n\tn(\n\t\tfunc(m *M) {\n\t\t\tm.H(\n\t\t\t\t\"/\",\n\t\t\t\tfunc(\n\t\t\t\t\tw W,\n\t\t\t\t\t_ R,\n\t\t\t\t) {\n\t\t\t\t\tw.Write()\n\t\t\t\t},\n\t\t\t)\n\t\t},\n\t)\n}\n\nfunc Second() {}\n",
		),
	)
	assertReport(
		t,
		"interfunc",
		false,
		nil,
		"package example\n\nfunc First() {\n\tn(\n\t\tfunc(m *M) {\n\t\t\tm.H(\n\t\t\t\t\"/\",\n\t\t\t\tfunc(\n\t\t\t\t\tw W,\n\t\t\t\t\t_ R,\n\t\t\t\t) {\n\t\t\t\t\tw.Write()\n\t\t\t\t},\n\t\t\t)\n\t\t},\n\t)\n}\n\nfunc Second() {}\n",
		l,
	)
}

func TestSpacingBlankBetweenTopLevelConstantAndVariablePreserved(t *testing.T) {
	input := "package example\n\nconst A = 1\n\nvar B = 2\n"
	l := Spacing("constvarblank", strings.NewReader(input))
	assertReport(t, "constvarblank", false, nil, input, l)
}

func TestSpacingBlankBetweenTopLevelVariableAndConstantPreserved(t *testing.T) {
	input := "package example\n\nvar A = 1\n\nconst B = 2\n"
	l := Spacing("varconstblank", strings.NewReader(input))
	assertReport(t, "varconstblank", false, nil, input, l)
}

func TestSpacingMissingBlankBetweenTopLevelConstantAndVariableInserted(t *testing.T) {
	l := Spacing(
		"constvarinsert",
		strings.NewReader("package example\n\nconst A = 1\nvar B = 2\n"),
	)
	assertReport(
		t,
		"constvarinsert",
		true,
		[]*concern.Concern{
			{
				Key:      constant.MissingBlankBetweenVariableConstantKey,
				Text:     constant.MissingBlankBetweenVariableConstantText,
				Path:     "constvarinsert",
				Type:     concern.Line,
				Line:     4,
				LineText: "var B = 2",
				Fixed:    true,
			},
		},
		"package example\n\nconst A = 1\n\nvar B = 2\n",
		l,
	)
}

func TestSpacingMissingBlankBetweenTopLevelVariableAndConstantInserted(t *testing.T) {
	l := Spacing(
		"varconstinsert",
		strings.NewReader("package example\n\nvar A = 1\nconst B = 2\n"),
	)
	assertReport(
		t,
		"varconstinsert",
		true,
		[]*concern.Concern{
			{
				Key:      constant.MissingBlankBetweenVariableConstantKey,
				Text:     constant.MissingBlankBetweenVariableConstantText,
				Path:     "varconstinsert",
				Type:     concern.Line,
				Line:     4,
				LineText: "const B = 2",
				Fixed:    true,
			},
		},
		"package example\n\nvar A = 1\n\nconst B = 2\n",
		l,
	)
}

func TestSpacingMultiLineConditionThreeLineBlankAfterControl(t *testing.T) {
	input := "package example\n\nfunc Example() {\n\tif a == 0 &&\n\t\tb == 0 &&\n\t\tc == 0 {\n\t\treturn\n\t}\n\n\tx := 1\n\t_ = x\n}\n"
	l := Spacing("multiline3", strings.NewReader(input))
	assertReport(t, "multiline3", false, nil, input, l)
}

func TestSpacingDoubleBlankAfterClosingBrace(t *testing.T) {
	l := Spacing(
		upper.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\n\tfmt.Println(\"b\")\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      constant.ExtraneousBlankLineKey,
				Text:     constant.ExtraneousBlankLineText,
				Path:     "Alfa",
				Type:     concern.Line,
				Line:     8,
				LineText: "",
				Fixed:    true,
			},
		},
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		l,
	)
}
