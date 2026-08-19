package store

import "gorm.io/gorm"

func revocationClause(
	d *gorm.DB,
	revoked *bool,
) *gorm.DB {
	if revoked == nil {
		return d
	}

	if *revoked {
		return d.Where("revoked_at IS NOT NULL")
	}

	return d.Where("revoked_at IS NULL")
}
