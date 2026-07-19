package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/text/markdown/file"
	"testing"
)

func TestMarkdownFileParse(t *testing.T) {
	s := "## Example\n\nList of files in the current directory:\n```sh\nls -alh\n```\n"
	f := file.New(new([]byte(s)))
	assert.Strings(
		t,
		[]string{
			"Example",
			"List of files in the current directory:",
			"ls -alh\n",
		},
		f.Parse().Content(),
	)
}

func TestFixture(t *testing.T) {
	assert.String(
		t,
		"## Example\n\nList of files in the current directory:\n```sh\nls -alh\n```\n",
		fixture.Read(constant.MarkdownPath, "1.md"),
	)
}
