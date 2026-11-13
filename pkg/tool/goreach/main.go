package goreach

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goreach/constant"
)

func Main() {
	s := instrument.NewCommandLine(constant.Identity)
	defer func() { s.Flush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		constant.Transcript,
		"",
		"Transcript to read instead of the session's, a .jsonl path",
	)
	a.Boolean(
		constant.Documents,
		false,
		"Also print the directories of changed files that are not Go, for golint",
	)
	a.Boolean(
		constant.Tags,
		false,
		"Print the repository's custom build tags comma-joined, for golangci-lint",
	)
	a.String(
		argumentConstant.Root,
		"",
		"Repository root, discovered upward from the working directory when empty",
	)
	a.Parse()
	root, _ := lint.Root(a.GetString(argumentConstant.Root))

	if a.GetBoolean(constant.Tags) {
		console.Format("%s\n", Tags(root))

		return
	}

	transcript, e := TranscriptPath(
		a.GetString(constant.Transcript),
		a.Positionals(),
	)

	if e != nil {
		system.Exitf(1, "%s\n", e)
	}

	Run(root, transcript, a.GetBoolean(constant.Documents))
}
