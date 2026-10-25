package statement

import "database/sql/driver"

func named(v []driver.Value) []driver.NamedValue {
	result := make([]driver.NamedValue, len(v))

	for i, value := range v {
		result[i] = driver.NamedValue{Ordinal: i + 1, Value: value}
	}

	return result
}
