//go:build local

package tracker

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
	"testing"
)

func writeLine(
	f *os.File,
	line string,
) {
	_, e := fmt.Fprintf(f, "%s\n", line)

	if e != nil {
		panic(e)
	}
}

func TestColdStart(t *testing.T) {
	d := t.TempDir()
	path := filepath.Join(d, "session.jsonl")
	f, e := os.Create(path)

	if e != nil {
		t.Fatal(e)
	}

	writeLine(f, `{"type":"permission-mode","sessionId":"abc"}`)
	writeLine(
		f,
		`{"type":"user","timestamp":"2026-05-26T11:00:00.000Z","sessionId":"abc","slug":"my-slug","cwd":"/home/user","gitBranch":"main","message":{"role":"user","content":"can you help me fix the login bug"}}`,
	)
	writeLine(
		f,
		`{"type":"assistant","timestamp":"2026-05-26T11:01:00.000Z","sessionId":"abc","message":{"role":"assistant","content":"hi"}}`,
	)
	writeLine(
		f,
		`{"type":"user","timestamp":"2026-05-26T12:00:00.000Z","sessionId":"abc","message":{"role":"user","content":"now update the tests to match"}}`,
	)
	errors.PanicClose(f)
	s := &State{}
	assert.Nil(t, Read(path, s))
	assert.Integer(t, 4, s.Lines)
	assert.String(t, "2026-05-26T11:00:00.000Z", s.FirstTimestamp)
	assert.String(t, "2026-05-26T12:00:00.000Z", s.LastTimestamp)
	assert.String(t, "my-slug", s.Slug)
	assert.String(t, "/home/user", s.WorkDirectory)
	assert.String(t, "main", s.Branch)
	assert.True(t, s.Offset > 0)
	assert.Integer(t, 2, s.UserMessageCount)
	assert.String(t, "can you help me fix the login bug", s.FirstMessage)
}

func TestIncremental(t *testing.T) {
	d := t.TempDir()
	path := filepath.Join(d, "session.jsonl")
	f, e := os.Create(path)

	if e != nil {
		t.Fatal(e)
	}

	writeLine(
		f,
		`{"type":"user","timestamp":"2026-05-26T11:00:00.000Z","sessionId":"abc","slug":"my-slug","cwd":"/home/user","gitBranch":"main","message":{"role":"user","content":"can you help me fix the login bug"}}`,
	)
	writeLine(
		f,
		`{"type":"assistant","timestamp":"2026-05-26T11:01:00.000Z","sessionId":"abc","message":{"role":"assistant","content":"hi"}}`,
	)
	errors.PanicClose(f)
	s := &State{}
	assert.Nil(t, Read(path, s))
	assert.Integer(t, 2, s.Lines)
	assert.Integer(t, 1, s.UserMessageCount)
	assert.String(t, "can you help me fix the login bug", s.FirstMessage)
	savedOffset := s.Offset
	f, e = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)

	if e != nil {
		t.Fatal(e)
	}

	writeLine(
		f,
		`{"type":"user","timestamp":"2026-05-26T12:00:00.000Z","sessionId":"abc","message":{"role":"user","content":"now update the tests to match"}}`,
	)
	writeLine(
		f,
		`{"type":"assistant","timestamp":"2026-05-26T12:01:00.000Z","sessionId":"abc","message":{"role":"assistant","content":"ok"}}`,
	)
	errors.PanicClose(f)
	assert.Nil(t, Read(path, s))
	assert.Integer(t, 4, s.Lines)
	assert.Integer(t, 2, s.UserMessageCount)
	assert.String(t, "can you help me fix the login bug", s.FirstMessage)
	assert.String(t, "2026-05-26T12:01:00.000Z", s.LastTimestamp)
	assert.String(t, "my-slug", s.Slug)
	assert.True(t, s.Offset > savedOffset)
}

func TestIncrementalDoesNotOverwriteColdFields(t *testing.T) {
	d := t.TempDir()
	path := filepath.Join(d, "session.jsonl")
	f, e := os.Create(path)

	if e != nil {
		t.Fatal(e)
	}

	writeLine(
		f,
		`{"type":"user","timestamp":"2026-05-26T11:00:00.000Z","sessionId":"abc","slug":"original","cwd":"/original","gitBranch":"main"}`,
	)
	errors.PanicClose(f)
	s := &State{}
	assert.Nil(t, Read(path, s))
	f, e = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)

	if e != nil {
		t.Fatal(e)
	}

	writeLine(
		f,
		`{"type":"user","timestamp":"2026-05-26T12:00:00.000Z","sessionId":"abc","slug":"different","cwd":"/different","gitBranch":"feature"}`,
	)
	errors.PanicClose(f)
	assert.Nil(t, Read(path, s))
	assert.String(t, "original", s.Slug)
	assert.String(t, "/original", s.WorkDirectory)
	assert.String(t, "main", s.Branch)
}

func TestMissingFile(t *testing.T) {
	s := &State{}
	e := Read("/nonexistent/path.jsonl", s)
	assert.True(t, e != nil)
	assert.Integer(t, 0, s.Lines)
}

func TestNoTimestamps(t *testing.T) {
	d := t.TempDir()
	path := filepath.Join(d, "session.jsonl")
	f, e := os.Create(path)

	if e != nil {
		t.Fatal(e)
	}

	writeLine(f, `{"type":"permission-mode","sessionId":"abc"}`)
	writeLine(f, `{"type":"file-history-snapshot"}`)
	errors.PanicClose(f)
	s := &State{}
	assert.Nil(t, Read(path, s))
	assert.Integer(t, 2, s.Lines)
	assert.String(t, "", s.FirstTimestamp)
	assert.String(t, "", s.LastTimestamp)
}
