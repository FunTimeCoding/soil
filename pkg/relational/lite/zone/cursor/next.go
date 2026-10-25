package cursor

import (
	"database/sql/driver"
	"time"
)

func (c *Cursor) Next(v []driver.Value) error {
	if e := c.inner.Next(v); e != nil {
		return e
	}

	for i, value := range v {
		if t, okay := value.(time.Time); okay {
			v[i] = t.Local()
		}
	}

	return nil
}
