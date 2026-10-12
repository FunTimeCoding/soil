package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestPackageNameBlacklisted(t *testing.T) {
	l := lint.PackageName(
		stringsConstant.UpperAlfa,
		strings.NewReader("package api\n"),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "package_name",
				Text:     "Blacklisted package name",
				Path:     "Alfa",
				Type:     constant.ConcernLine,
				Line:     1,
				LineText: "package api",
			},
		},
		"",
		l,
	)
}

func TestPackageNameAllowed(t *testing.T) {
	l := lint.PackageName(
		stringsConstant.UpperBravo,
		strings.NewReader("package server\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}
