package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"os"
	"path/filepath"
	"testing"
)

func shopRepository(t *testing.T) *repository.Repository {
	t.Helper()
	directory := t.TempDir()
	v := virtual_file_system.New()

	for path, content := range map[string]string{
		"pkg/shop/shop.go":           "package shop\n\nimport \"sync\"\n\ntype Base struct{ Name string }\n\nfunc (b *Base) Label() string { return b.Name }\n\ntype Item struct {\n\tBase\n\tPrice int\n}\n\ntype Locked struct{ sync.Mutex }\n\nvar Default = &Item{}\n",
		"pkg/shop/unit/shop_test.go": "package unit\n\nimport \"testing\"\n\nfunc TestShopKeepsStock(t *testing.T) {}\n",
		"pkg/stock/unit/stock_test.go": "package unit\n\nimport \"testing\"\n\nfunc TestShopKeepsStock(t *testing.T) {}\n\nfunc TestStockCounts(t *testing.T) {}\n",
	} {
		full := filepath.Join(directory, path)
		assert.FatalOnError(t, os.MkdirAll(filepath.Dir(full), 0o755))
		assert.FatalOnError(t, os.WriteFile(full, []byte(content), 0o644))
		v.WriteString(path, content)
	}

	return repository.New(directory, v)
}
