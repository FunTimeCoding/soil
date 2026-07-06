package lint

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestIsExecutableElf(t *testing.T) {
	assert.Boolean(t, true, isExecutable("\x7fELF\x02\x01\x01"))
}

func TestIsExecutableMachO64(t *testing.T) {
	assert.Boolean(t, true, isExecutable("\xcf\xfa\xed\xfe\x07\x00"))
}

func TestIsExecutableMachO32(t *testing.T) {
	assert.Boolean(t, true, isExecutable("\xce\xfa\xed\xfe\x07\x00"))
}

func TestIsExecutableMachOBig(t *testing.T) {
	assert.Boolean(t, true, isExecutable("\xfe\xed\xfa\xcf\x00\x00"))
}

func TestIsExecutablePE(t *testing.T) {
	assert.Boolean(t, true, isExecutable("MZ\x90\x00"))
}

func TestIsExecutableWasm(t *testing.T) {
	assert.Boolean(t, true, isExecutable("\x00asm\x01\x00"))
}

func TestIsExecutableGoSource(t *testing.T) {
	assert.Boolean(t, false, isExecutable("package main\n"))
}

func TestIsExecutableEmpty(t *testing.T) {
	assert.Boolean(t, false, isExecutable(""))
}

func TestIsExecutableTooShort(t *testing.T) {
	assert.Boolean(t, false, isExecutable("abc"))
}
