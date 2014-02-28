package ω

import (
	"fmt"
	"reflect"
)

type Value interface{}
type Slice []interface{}

var iv []interface{}
var tt = reflect.TypeOf(iv)

func ToString(v interface{}) string {
	return fmt.Sprint(v)
}

func I(v Value) Value {
	return fmt.Sprint(v)
}

func S(v Value) Value {
	return fmt.Sprint(v)
}

func A(v interface{}) Slice {
	rv := reflect.ValueOf(v)
	la := rv.Len()
	t := reflect.MakeSlice(tt, la, la)
	for i := 0; i < la; i++ {
		t.Index(i).Set(rv.Index(i))
	}
	return t.Interface().([]interface{})
}

func (s Slice) C(v interface{}) interface{} {
	rs := reflect.ValueOf(s)
	la := reflect.ValueOf(s).Len()
	t := reflect.MakeSlice(reflect.TypeOf(v), la, la)
	for i := 0; i < la; i++ {
		se := rs.Index(i).Elem()
		te := t.Index(i)
		if se.Type().ConvertibleTo(te.Type()) {
			te.Set(se.Convert(te.Type()))
		} else {
			te.Set(se)
		}
	}
	return t.Interface()
}

func (s Slice) N() interface{} {
	if len(s) == 0 {
		return s
	}
	rs := reflect.ValueOf(s)
	la := reflect.ValueOf(s).Len()
	t := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(s[0])), la, la)
	for i := 0; i < la; i++ {
		se := rs.Index(i).Elem()
		te := t.Index(i)
		if se.Type().ConvertibleTo(te.Type()) {
			te.Set(se.Convert(te.Type()))
		} else {
			te.Set(se)
		}
	}
	return t.Interface()
}

func (s Slice) Map(f func(v Value) Value) Slice {
	rs := reflect.ValueOf(s)
	ls := rs.Len()
	rf := reflect.ValueOf(f)
	t := reflect.MakeSlice(tt, ls, ls)
	for i := 0; i < ls; i++ {
		ret := rf.Call([]reflect.Value{rs.Index(i)})
		t.Index(i).Set(ret[0])
	}
	return t.Interface().([]interface{})
}

func (s Slice) ForceMap(f func(v Value) Value) Slice {
	rs := reflect.ValueOf(s)
	ls := rs.Len()
	rf := reflect.ValueOf(f)
	for i := 0; i < ls; i++ {
		ret := rf.Call([]reflect.Value{rs.Index(i)})
		rs.Index(i).Set(ret[0])
	}
	return rs.Interface().([]interface{})
}

func (s Slice) Filter(f func(v Value) bool) Slice {
	rs := reflect.ValueOf(s)
	ls := rs.Len()
	rf := reflect.ValueOf(f)
	t := reflect.MakeSlice(tt, 0, ls)
	for i := 0; i < ls; i++ {
		ret := rf.Call([]reflect.Value{rs.Index(i)})
		if ret[0].Interface().(bool) {
			t = reflect.Append(t, rs.Index(i))
		}
	}
	return t.Interface().([]interface{})
}

func (s Slice) Each(f func(v Value)) Slice {
	rs := reflect.ValueOf(s)
	ls := rs.Len()
	rf := reflect.ValueOf(f)
	for i := 0; i < ls; i++ {
		rf.Call([]reflect.Value{rs.Index(i)})
	}
	return s
}
