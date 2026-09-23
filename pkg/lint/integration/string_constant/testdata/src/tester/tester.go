package tester

import (
	"tester.test/constant"
	"tester.test/widget_tester"
	"testing"
)

func MethodClean(
	t *testing.T,
	actual string,
) {
	widget_tester.New(t).AssertName("name", actual)
}

func MethodTautology(
	t *testing.T,
	actual string,
) {
	widget_tester.New(t).AssertName(
		constant.Name,
		actual,
	) // want `constant constant.Name in expected value should be a literal`
}

func InputFlagged(t *testing.T) string {
	return widget_tester.New(t).Submit(
		t,
		"name",
	) // want `string literal "name" has constant constant.Name`
}
