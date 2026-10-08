package chromium

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/notation"
	"reflect"
)

func Decode(
	value []byte,
	result any,
) error {
	if result == nil {
		return nil
	}

	if value != nil {
		return notation.DecodeBytes(value, result)
	}

	v := reflect.ValueOf(result).Elem()

	switch v.Kind() {
	case reflect.Pointer,
		reflect.Map,
		reflect.Slice,
		reflect.Chan,
		reflect.Func,
		reflect.Interface:
		v.SetZero()

		return nil
	default:
		return chromedp.ErrJSUndefined
	}
}
