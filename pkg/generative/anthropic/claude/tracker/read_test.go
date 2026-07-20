package tracker

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"testing"
)

func assistantLine(
	messageIdentifier string,
	request string,
	model string,
	usage string,
) string {
	return join.Empty(
		fmt.Sprintf(
			`{"type":"assistant","timestamp":"2026-07-20T10:00:00.000Z","requestId":%q,"message":{"id":%q,"role":"assistant","model":%q,"usage":%s}}`,
			request,
			messageIdentifier,
			model,
			usage,
		),
		"\n",
	)
}

func appendFile(
	t *testing.T,
	path string,
	content string,
) {
	f, e := os.OpenFile(
		path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if e != nil {
		t.Fatalf("open %s: %v", path, e)
	}

	if _, e = f.WriteString(content); e != nil {
		t.Fatalf("write %s: %v", path, e)
	}

	if e = f.Close(); e != nil {
		t.Fatalf("close %s: %v", path, e)
	}
}

func TestReadAccumulatesUsage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "session.jsonl")
	fable := `{"input_tokens":10,"output_tokens":100,"cache_creation_input_tokens":1000,"cache_read_input_tokens":5000,"cache_creation":{"ephemeral_5m_input_tokens":200,"ephemeral_1h_input_tokens":800}}`
	appendFile(t, path, assistantLine("m1", "r1", "claude-fable-5", fable))
	appendFile(t, path, assistantLine("m1", "r1", "claude-fable-5", fable))
	s := New()
	assert.Nil(t, Read(path, s))
	tokens := s.Usage["fable"]
	assert.NotNil(t, tokens)
	assert.Integer(t, 1, tokens.Count)
	assert.Integer(t, 10, tokens.Input)
	assert.Integer(t, 100, tokens.Output)
	assert.Integer(t, 1000, tokens.CacheCreation)
	assert.Integer(t, 200, tokens.Cache5Minute)
	assert.Integer(t, 800, tokens.Cache1Hour)
	assert.Integer(t, 5000, tokens.CacheRead)
	appendFile(t, path, assistantLine("m1", "r1", "claude-fable-5", fable))
	opus := `{"input_tokens":7,"output_tokens":30,"cache_creation_input_tokens":400,"cache_read_input_tokens":900}`
	appendFile(t, path, assistantLine("m2", "r2", "claude-opus-4-8", opus))
	assert.Nil(t, Read(path, s))
	tokens = s.Usage["fable"]
	assert.Integer(t, 1, tokens.Count)
	assert.Integer(t, 10, tokens.Input)
	second := s.Usage["opus"]
	assert.NotNil(t, second)
	assert.Integer(t, 1, second.Count)
	assert.Integer(t, 7, second.Input)
	assert.Integer(t, 400, second.CacheCreation)
	assert.Integer(t, 0, second.Cache5Minute)
	assert.Integer(t, 0, second.Cache1Hour)
}

func TestReadResetsOnShrink(t *testing.T) {
	path := filepath.Join(t.TempDir(), "session.jsonl")
	fable := `{"input_tokens":10,"output_tokens":100,"cache_creation_input_tokens":0,"cache_read_input_tokens":0}`
	appendFile(t, path, assistantLine("m1", "r1", "claude-fable-5", fable))
	appendFile(t, path, assistantLine("m2", "r2", "claude-fable-5", fable))
	s := New()
	assert.Nil(t, Read(path, s))
	assert.Integer(t, 2, s.Usage["fable"].Count)

	if e := os.WriteFile(
		path,
		[]byte(assistantLine("m3", "r3", "claude-fable-5", fable)),
		0644,
	); e != nil {
		t.Fatalf("rewrite %s: %v", path, e)
	}

	assert.Nil(t, Read(path, s))
	assert.Integer(t, 1, s.Usage["fable"].Count)
	assert.Integer(t, 1, s.Lines)
}
