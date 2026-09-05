package unit

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestPackageNameBlacklisted(t *testing.T) {
	l := lint.PackageName(
		constant.UpperAlfa,
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
				Type:     lintConstant.ConcernLine,
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
		constant.UpperBravo,
		strings.NewReader("package server\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}
