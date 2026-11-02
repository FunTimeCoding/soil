package reach

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"os"
	"slices"
	"sort"
)

func ReadTranscript(
	path string,
	modules map[string]string,
) (*Edits, error) {
	f, e := os.Open(path)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(f)
	files := map[string]bool{}
	pending := map[string]bool{}
	module := ""
	result := NewEdits(nil)
	s := bufio.NewScanner(f)
	s.Buffer(
		make([]byte, constant.TranscriptLineLimit),
		constant.TranscriptLineLimit,
	)

	for s.Scan() {
		cwd, found := blocks(s.Bytes())

		for _, b := range found {
			switch {
			case b.Type == constant.TranscriptToolUse &&
				slices.Contains(constant.EditTools, b.Name):

				if p := inputString(b.Input, constant.FilePathInput); p != "" {
					files[p] = true
				}
			case b.Type == constant.TranscriptToolUse && b.Name == constant.SourceUseModule:
				module = inputString(b.Input, constant.ModuleInput)
			case b.Type == constant.TranscriptToolUse && isSourceWrite(b):
				pending[b.Identifier] = true
			case b.Type == constant.TranscriptToolUse && b.Name == constant.BashTool:
				command := inputString(b.Input, constant.CommandInput)
				addCommand(result, files, ResolveCommand(cwd, command))
			case b.Type == constant.TranscriptToolResult && pending[b.ToolUse]:
				for _, p := range resultPaths(b.Content) {
					files[resolvePath(modules, module, p)] = true
				}
			}
		}
	}

	if f := s.Err(); f != nil {
		return nil, f
	}

	for p := range files {
		result.Files = append(result.Files, p)
	}

	sort.Strings(result.Files)

	return result, nil
}
