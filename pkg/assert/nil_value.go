package assert

import "reflect"

func nilValue(actual any) bool {
	if actual == nil {
		return true
	}

	v := reflect.ValueOf(actual)

	switch v.Kind() {
	case reflect.Pointer,
		reflect.Slice,
		reflect.Map,
		reflect.Chan,
		reflect.Func,
		reflect.Interface:

		return v.IsNil()
	default:
		return false
	}
}
