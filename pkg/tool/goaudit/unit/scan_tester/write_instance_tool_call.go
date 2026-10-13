package scan_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func WriteInstanceToolCall(base string) {
	system.WriteFile(
		filepath.Join(base, "pkg/tool/goalfad/model_context/new.go"),
		[]byte(
			"package model_context\n\nfunc New() {\n\tinstance_tool.New(nil, nil)\n}\n",
		),
		0644,
	)
}
