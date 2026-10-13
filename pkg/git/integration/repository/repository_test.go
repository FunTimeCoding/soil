package repository

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git"
	"testing"
	"time"
)

func TestCommitTimes(t *testing.T) {
	times, e := git.CommitTimes(git.FindDirectory())
	assert.FatalOnError(t, e)
	modified, okay := times["go.mod"]
	assert.True(t, okay)
	assert.True(t, modified.After(time.Unix(0, 0)))
}

func TestModifiedFiles(t *testing.T) {
	_, e := git.ModifiedFiles(git.FindDirectory())
	assert.FatalOnError(t, e)
}
