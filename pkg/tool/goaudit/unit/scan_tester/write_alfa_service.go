package scan_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func WriteAlfaService(
	base string,
	settings string,
) {
	system.MakeDirectory(filepath.Join(base, ".claude"))
	system.WriteFile(
		filepath.Join(base, ".claude", "settings.local.json"),
		[]byte(settings),
		0644,
	)
	system.MakeDirectory(filepath.Join(base, "pkg/tool/goalfad/constant"))
	system.MakeDirectory(filepath.Join(base, "pkg/tool/goalfad/model_context"))
	system.WriteFile(
		filepath.Join(base, "pkg/tool/goalfad/constant/constant.go"),
		[]byte("package constant\n\nconst AlfaSearch = \"alfa_search\"\n"),
		0644,
	)
	system.WriteFile(
		filepath.Join(base, "pkg/tool/goalfad/model_context/register.go"),
		[]byte(
			"package model_context\n\nfunc register() {\n\tadd(mcp.NewTool(constant.AlfaSearch))\n}\n",
		),
		0644,
	)
}
