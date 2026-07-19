package client

import (
	"fmt"
	"reflect"
)

// Validate reports the first unset field, so constructor wiring
// stays covered even for fields added after the check was written.
func (c *Client) Validate() error {
	v := reflect.ValueOf(*c)

	for i := range v.NumField() {
		f := v.Field(i)
		name := v.Type().Field(i).Name

		switch f.Kind() {
		case reflect.Pointer, reflect.Interface, reflect.Map:
			if f.IsNil() {
				return fmt.Errorf("field not set: %s", name)
			}
		case reflect.String:
			if f.String() == "" {
				return fmt.Errorf("field empty: %s", name)
			}
		default:
			continue
		}
	}

	return nil
}
