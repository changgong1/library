package util

import (
	"reflect"
)

func SafeGo(f interface{}, args ...interface{}) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				// logger.Error("SafeGo ", zap.Any("f", f), zap.Any("param", args), zap.Any("err", err), zap.Any("stack", string(debug.Stack())))
			}
		}()

		v1 := reflect.ValueOf(f)
		params := make([]reflect.Value, len(args))
		for k, v := range args {
			params[k] = reflect.ValueOf(v)
		}
		v1.Call(params)
	}()
}
