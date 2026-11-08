package gorm

type DB struct{}

func (d *DB) Where(
	query any,
	arguments ...any,
) *DB {
	return d
}

func (d *DB) First(
	destination any,
	conditions ...any,
) *DB {
	return d
}

func (d *DB) Save(value any) *DB {
	return d
}

func (d *DB) Model(value any) *DB {
	return d
}

func (d *DB) Scan(destination any) *DB {
	return d
}
