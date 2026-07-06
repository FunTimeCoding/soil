package goarchive

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/goarchive/constant"
	"path/filepath"
	"strings"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		"output",
		"",
		"output directory (default: strip archive extension)",
	)
	a.String("password", "", "archive password")
	a.Parse(version, gitHash, buildDate)
	path := a.RequiredPositional(0, "file")
	output := a.GetString("output")
	password := a.GetString("password")

	if output == "" {
		output = strings.TrimSuffix(path, filepath.Ext(path))
	}

	extract(path, output, password)
}
