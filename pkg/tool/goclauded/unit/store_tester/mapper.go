package store_tester

import "gorm.io/gorm"

func (t *Tester) Mapper() *gorm.DB {
	return t.mapper
}
