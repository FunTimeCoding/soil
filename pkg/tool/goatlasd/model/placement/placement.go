package placement

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/attribution"

type Placement struct {
	Identifier              uint   `gorm:"primaryKey;column:identifier"`
	Source                  string `gorm:"column:source"`
	Kind                    string `gorm:"column:kind"`
	Scope                   string `gorm:"column:scope"`
	Name                    string `gorm:"column:name"`
	Package                 string `gorm:"column:package"`
	Version                 string `gorm:"column:version"`
	attribution.Attribution `gorm:"embedded"`
}
