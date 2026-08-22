package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestRegisteredToolsResolvesConstants(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"constant/constant.go",
		"package constant\n\nconst (\n\tAlfaSearch = \"alfa_search\"\n\tAlfaPost   = \"alfa_post\"\n)\n",
	)
	v.WriteString(
		"model_context/register.go",
		"package model_context\n\nfunc register() {\n\tadd(mcp.NewTool(constant.AlfaSearch))\n\tadd(mcp.NewTool(constant.AlfaPost))\n\tadd(mcp.NewTool(\"alfa_direct\"))\n}\n",
	)
	names, concerns := scan.RegisteredTools(v, "pkg/tool/goalfad")
	assert.Integer(t, 0, len(concerns))
	assert.Integer(t, 3, len(names))
	assert.String(t, "alfa_search", names[0])
	assert.String(t, "alfa_post", names[1])
	assert.String(t, "alfa_direct", names[2])
}

func TestRegisteredToolsSplitFiles(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"constant/constant.go",
		"package constant\n\nconst AlfaList = \"alfa_list\"\n",
	)
	v.WriteString(
		"model_context/alfa_list_tool.go",
		"package model_context\n\nfunc AlfaListTool() mcp.Tool {\n\treturn mcp.NewTool(constant.AlfaList)\n}\n",
	)
	names, concerns := scan.RegisteredTools(v, "pkg/tool/goalfad")
	assert.Integer(t, 0, len(concerns))
	assert.Integer(t, 1, len(names))
	assert.String(t, "alfa_list", names[0])
}

func TestRegisteredToolsSubpackage(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"constant/constant.go",
		"package constant\n\nconst AlfaAdd = \"alfa_add\"\n",
	)
	v.WriteString(
		"model_context/tool/add.go",
		"package tool\n\nfunc Add() mcp.Tool {\n\treturn mcp.NewTool(constant.AlfaAdd)\n}\n",
	)
	names, concerns := scan.RegisteredTools(v, "pkg/tool/goalfad")
	assert.Integer(t, 0, len(concerns))
	assert.Integer(t, 1, len(names))
	assert.String(t, "alfa_add", names[0])
}

func TestRegisteredToolsMissingModelContext(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("run.go", "package goalfad\n")
	names, concerns := scan.RegisteredTools(v, "pkg/tool/goalfad")
	assert.Integer(t, 0, len(names))
	assert.Integer(t, 1, len(concerns))
	assert.String(t, "mapped_service_missing", concerns[0].Key)
	assert.String(t, "pkg/tool/goalfad", concerns[0].Path)
}

func TestRegisteredToolsUnresolvedName(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"model_context/register.go",
		"package model_context\n\nfunc register() {\n\tadd(mcp.NewTool(constant.Unknown))\n}\n",
	)
	names, concerns := scan.RegisteredTools(v, "pkg/tool/goalfad")
	assert.Integer(t, 0, len(names))
	assert.Integer(t, 1, len(concerns))
	assert.String(t, "unresolved_tool_name", concerns[0].Key)
	assert.String(
		t,
		"pkg/tool/goalfad/model_context/register.go",
		concerns[0].Path,
	)
}
