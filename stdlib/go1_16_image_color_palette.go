// Code generated by 'yaegi extract image/color/palette'. DO NOT EDIT.

// +build go1.16

package stdlib

import (
	"image/color/palette"
	"reflect"
)

func init() {
	Symbols["image/color/palette"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Plan9":   reflect.ValueOf(&palette.Plan9).Elem(),
		"WebSafe": reflect.ValueOf(&palette.WebSafe).Elem(),
	}
}