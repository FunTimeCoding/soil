package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"testing"
)

func TestIsExecutableElf(t *testing.T) {
	assert.Boolean(t, true, lint.IsExecutable("\x7fELF\x02\x01\x01"))
}

func TestIsExecutableMachO64(t *testing.T) {
	assert.Boolean(t, true, lint.IsExecutable("\xcf\xfa\xed\xfe\x07\x00"))
}

func TestIsExecutableMachO32(t *testing.T) {
	assert.Boolean(t, true, lint.IsExecutable("\xce\xfa\xed\xfe\x07\x00"))
}

func TestIsExecutableMachOBig(t *testing.T) {
	assert.Boolean(t, true, lint.IsExecutable("\xfe\xed\xfa\xcf\x00\x00"))
}

func TestIsExecutablePE(t *testing.T) {
	assert.Boolean(t, true, lint.IsExecutable("MZ\x90\x00"))
}

func TestIsExecutableWasm(t *testing.T) {
	assert.Boolean(t, true, lint.IsExecutable("\x00asm\x01\x00"))
}

func TestIsExecutableGoSource(t *testing.T) {
	assert.Boolean(t, false, lint.IsExecutable("package main\n"))
}

func TestIsExecutableEmpty(t *testing.T) {
	assert.Boolean(t, false, lint.IsExecutable(""))
}

func TestIsExecutableTooShort(t *testing.T) {
	assert.Boolean(t, false, lint.IsExecutable("abc"))
}
