package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"gorm.io/gorm/schema"
	"sync"
	"testing"
)

func TestFreshStoreCarriesEveryModelColumnAndIndex(t *testing.T) {
	d := store_tester.New(t).Store.Database()
	var missing []string

	for _, model := range migrate.Models() {
		parsed, e := schema.Parse(model, &sync.Map{}, d.NamingStrategy)
		assert.FatalOnError(t, e)

		for _, column := range parsed.DBNames {
			if !d.Migrator().HasColumn(model, column) {
				missing = append(
					missing,
					join.Empty(parsed.Table, stringsConstant.Dot, column),
				)
			}
		}

		for _, index := range parsed.ParseIndexes() {
			if !d.Migrator().HasIndex(model, index.Name) {
				missing = append(
					missing,
					join.Empty(parsed.Table, " index ", index.Name),
				)
			}
		}
	}

	assert.Strings(t, nil, missing)
}
