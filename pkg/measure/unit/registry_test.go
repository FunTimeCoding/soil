package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/measure/language"
	"github.com/funtimecoding/soil/pkg/measure/registry"
	"testing"
)

func TestByPathSuffixBeatsExtension(t *testing.T) {
	r := registry.NewDefault()
	assert.String(t, "Go Test", r.ByPath("pkg/alfa/bravo_test.go").Name)
	assert.String(t, "Go", r.ByPath("pkg/alfa/bravo.go").Name)
}

func TestByPathFilename(t *testing.T) {
	r := registry.NewDefault()
	assert.String(t, "Makefile", r.ByPath("charlie/Makefile").Name)
	assert.String(t, "Makefile", r.ByPath("charlie/rules.mk").Name)
}

func TestByPathUnknown(t *testing.T) {
	r := registry.NewDefault()
	assert.Nil(t, r.ByPath("delta.unknown"))
	assert.Nil(t, r.ByPath("noextension"))
}

func TestByShebang(t *testing.T) {
	r := registry.New(
		language.New("Shell").WithShebang("sh", "bash"),
		language.New("Python").WithShebang("python"),
	)
	assert.String(t, "Shell", r.ByShebang("#!/bin/sh").Name)
	assert.String(t, "Shell", r.ByShebang("#! /usr/bin/env bash").Name)
	assert.String(t, "Python", r.ByShebang("#!/usr/bin/env python3.12").Name)
	assert.Nil(t, r.ByShebang("#!/usr/bin/env"))
	assert.Nil(t, r.ByShebang("package main"))
}

func TestByName(t *testing.T) {
	r := registry.NewDefault()
	assert.String(t, "Go Test", r.ByName("go test").Name)
	assert.Nil(t, r.ByName("cobol"))
}
