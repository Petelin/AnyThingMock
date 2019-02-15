package mock

import "reflect"

type Mock struct {
	innerRec []func()
}

func (m *Mock) On(item interface{}, new interface{}) func() {
	old := reflect.ValueOf(item).Elem().Interface()
	rec := func() {
		if old == nil {
			v := reflect.New(reflect.TypeOf(item).Elem())
			reflect.ValueOf(item).Elem().Set(v.Elem())
		} else {
			reflect.ValueOf(item).Elem().Set(reflect.ValueOf(old))
		}
	}
	reflect.ValueOf(item).Elem().Set(reflect.ValueOf(new))
	m.innerRec = append(m.innerRec, rec)
	return rec
}

func (m *Mock) Recover() {
	for _, f := range m.innerRec {
		f()
	}
}
