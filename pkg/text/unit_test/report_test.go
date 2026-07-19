package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/report"
	"testing"
)

func TestReport(t *testing.T) {
	root := report.New("Example root", report.NoLimit)
	root.String("String", "example")
	root.Integer("Integer", 2)
	firstSection := root.Nest("Example first section", report.NoLimit)
	firstSection.String("String", "example")
	firstSection.Float("Float without unit", 1, "")
	firstSection.Float("Float with unit", 2, "L")
	firstSection.Integer("Integer", 1)
	firstSection.Percent("Percent", 50)
	secondSection := firstSection.Nest("Example second section", report.NoLimit)
	secondSection.String("String", "example")
	other := report.New("Example other", report.NoLimit)
	other.String("String", "other")
	root.AppendSection(other)
	assert.String(
		t,
		"Example root\n  String: example\n  Integer: 2\n  Example first section\n    String: example\n    Float without unit: 1.0\n    Float with unit: 2.0 L\n    Integer: 1\n    Percent: 50%\n    Example second section\n      String: example\n  Example other\n    String: other",
		root.Render(),
	)
}

func TestReportLimit(t *testing.T) {
	root := report.New("Example root", 67)
	assert.Integer(t, 12, root.Length())
	root.String("String", "example")
	assert.Integer(t, 30, root.Length())
	firstSection := root.Nest("Example section", report.NoLimit)
	assert.Integer(t, 15, firstSection.Length())
	firstSection.String("String", "example")
	assert.Integer(t, 37, firstSection.Length())
	secondSection := root.Nest("Too long section", report.NoLimit)
	assert.Integer(t, 16, secondSection.Length())
	assert.String(
		t,
		"Example root\n  String: example\n  Example section\n    String: example",
		root.Render(),
	)
}
