package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"strings"
	"testing"
)

func TestPackageNameBlacklisted(t *testing.T) {
	l := PackageName(
		upper.Alfa,
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
				Type:     concern.Line,
				Line:     1,
				LineText: "package api",
			},
		},
		"",
		l,
	)
}

func TestPackageNameAllowed(t *testing.T) {
	l := PackageName(
		upper.Bravo,
		strings.NewReader("package server\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}
