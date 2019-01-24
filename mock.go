package Mock

type Mock struct {
	innerRec []func()
}

func (m *Mock) On(item interface{}, new interface{}) *Mock {
	old := reflect.ValueOf(item).Elem().Interface()
	rec := func() {
		reflect.ValueOf(item).Elem().Set(reflect.ValueOf(old))
	}
	reflect.ValueOf(item).Elem().Set(reflect.ValueOf(new))
	m.innerRec = append(m.innerRec, rec)
	return m
}

func (m *Mock) Recover(item interface{}, new interface{}) {
	for _, f := range m.innerRec {
		f()
	}
}
