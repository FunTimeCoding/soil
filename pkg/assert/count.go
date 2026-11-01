package assert

import (
	"reflect"
	"testing"
)

func Count(
	t *testing.T,
	expect int,
	countable any,
) {
	t.Helper()
	var actual int

	switch reflect.TypeOf(countable).Kind() {
	case reflect.Slice:
		actual = reflect.ValueOf(countable).Len()
	case reflect.Array:
		actual = reflect.ValueOf(countable).Len()
	case reflect.Map:
		actual = reflect.ValueOf(countable).Len()
	case reflect.Chan:
		actual = reflect.ValueOf(countable).Len()
	default:
		t.Errorf("Not countable: %#v", countable)
	}

	Integer(t, expect, actual)
}
