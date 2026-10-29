//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/shared_reranker"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	shared_reranker.Release()
	os.Exit(code)
}
