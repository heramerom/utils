package maps

import "reflect"

func Default(mapLike interface{}, key interface{}, def interface{}) interface{} {
	m := reflect.ValueOf(mapLike)
	if m.Kind() == reflect.Map {
		if reflect.TypeOf(key) != m.Type().Key() && m.Type().Key().Kind() != reflect.Interface {
			return def
		}
		v := m.MapIndex(reflect.ValueOf(key))
		if v.IsValid() {
			return v.Interface()
		} else {
			return def
		}
	}
	return def
}
