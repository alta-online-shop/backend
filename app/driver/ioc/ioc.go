package ioc

import (
	"reflect"

	"github.com/gowok/ioc"
)

func Bind(name interface{}, fn ioc.BindFunc) {
	t := reflect.TypeOf(name)
	ioc.Bind(ioc.Key(t.Name()), fn)
}

func Use(name interface{}) interface{} {
	t := reflect.TypeOf(name)
	val, _ := ioc.Use(ioc.Key(t.Name()))
	return val
}
