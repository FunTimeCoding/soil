package connection

import (
	"database/sql/driver"
	"time"
)

func (c *Connection) CheckNamedValue(v *driver.NamedValue) error {
	converted, e := driver.DefaultParameterConverter.ConvertValue(v.Value)

	if e != nil {
		return e
	}

	if t, okay := converted.(time.Time); okay {
		converted = t.UTC()
	}

	v.Value = converted

	return nil
}
