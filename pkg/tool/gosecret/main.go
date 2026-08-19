package gosecret

import (
	"github.com/funtimecoding/soil/pkg/argument"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	var (
		checkMode  bool
		encodeMode bool
		directory  string
	)
	a := argument.NewInstance(constant.Identity)
	a.BooleanVariable(
		&checkMode,
		"check",
		false,
		"Check mode: verify decoded files match secrets, exit non-zero if any mismatch",
	)
	a.BooleanVariable(
		&encodeMode,
		"encode",
		false,
		"Encode mode: write decoded file contents back into secret manifests",
	)
	a.StringVariable(
		&directory,
		"directory",
		library.CurrentDirectory,
		"Directory to scan for secret manifests",
	)
	a.Parse(version, gitHash, buildDate)

	if checkMode && encodeMode {
		errors.Printf("Error: --check and --encode are mutually exclusive\n")
		os.Exit(1)
	}

	mode := constant.Decode

	if checkMode {
		mode = constant.Check
	}

	if encodeMode {
		mode = constant.Encode
	}

	if e := Run(directory, mode); e != nil {
		errors.Printf("Error: %v\n", e)
		os.Exit(1)
	}
}
