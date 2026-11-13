package gobreak

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/test_event"
	"github.com/funtimecoding/soil/pkg/test_event/test_run"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
)

func test(
	root string,
	f *break_file.File,
	flags ...string,
) *test_run.Run {
	arguments := []string{
		constant.GoCommand,
		constant.TestCommand,
		constant.NotationFlag,
		constant.CountFlag,
	}

	if len(f.Tags) > 0 {
		arguments = append(arguments, constant.TagsFlag, join.Comma(f.Tags))
	}

	if f.Tests != "" {
		arguments = append(arguments, constant.RunFlag, f.Tests)
	}

	arguments = append(arguments, flags...)
	r := run.New().NoPanic()
	r.Directory = root

	return test_run.New(
		test_event.Parse(r.Start(append(arguments, f.Packages...)...)),
	)
}
