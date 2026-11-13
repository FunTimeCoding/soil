package relational

import "gorm.io/gorm"

func FindOne(
	q *gorm.DB,
	out any,
) (bool, error) {
	r := q.Limit(1).Find(out)

	return r.RowsAffected > 0, r.Error
}
