package map2sorted_sets

import (
	"reflect"
	"sort"
)

func SaveInt64(m map[int64][]int64, k int64, v int64) {
	slice, ok := m[k]
	if !ok {
		m[k] = []int64{v}
		return
	}
	i := sort.Search(len(slice), func(i int) bool { return slice[i] >= v })
	if i >= len(slice) {
		slice = append(slice, v)
		m[k] = slice
		return
	} else if slice[i] != v {
		newSlice := []int64{}
		newSlice = append(newSlice, slice[:i]...)
		newSlice = append(newSlice, v)
		newSlice = append(newSlice, slice[i:]...)
		m[k] = newSlice
	}
}

func SaveStructByIntField(m, k interface{}, v interface{}, field string) {
	mapValue, key, value := reflect.ValueOf(m), reflect.ValueOf(k), reflect.ValueOf(v)
	slice := mapValue.MapIndex(key)
	if !slice.IsValid() {
		slice = reflect.MakeSlice(mapValue.Type().Elem(), 0, 1)
		mapValue.SetMapIndex(key, reflect.Append(slice, value))
		return
	}
	target := value.FieldByName(field).Int()
	i := sort.Search(slice.Len(), func(i int) bool {
		return slice.Index(i).FieldByName(field).Int() >= target
	})
	if i >= slice.Len() {
		slice = reflect.Append(slice, value)
		mapValue.SetMapIndex(key, slice)
		return
	} else if slice.Index(i).FieldByName(field).Int() != target {
		newSlice := reflect.MakeSlice(mapValue.Type().Elem(), 0, 0)
		newSlice = reflect.AppendSlice(newSlice, slice.Slice(0, i))
		newSlice = reflect.Append(newSlice, value)
		newSlice = reflect.AppendSlice(newSlice, slice.Slice(i, slice.Len()))
		mapValue.SetMapIndex(key, newSlice)
		return
	}
	slice.Index(i).Set(value)
}

func SaveStructByStringField(m, k interface{}, v interface{}, field string) {
	mapValue, key, value := reflect.ValueOf(m), reflect.ValueOf(k), reflect.ValueOf(v)
	slice := mapValue.MapIndex(key)
	if !slice.IsValid() {
		slice = reflect.MakeSlice(mapValue.Type().Elem(), 0, 1)
		mapValue.SetMapIndex(key, reflect.Append(slice, value))
		return
	}
	target := value.FieldByName(field).String()
	i := sort.Search(slice.Len(), func(i int) bool {
		return slice.Index(i).FieldByName(field).String() >= target
	})
	if i >= slice.Len() {
		slice = reflect.Append(slice, value)
		mapValue.SetMapIndex(key, slice)
		return
	} else if slice.Index(i).FieldByName(field).String() != target {
		newSlice := reflect.MakeSlice(mapValue.Type().Elem(), 0, 0)
		newSlice = reflect.AppendSlice(newSlice, slice.Slice(0, i))
		newSlice = reflect.Append(newSlice, value)
		newSlice = reflect.AppendSlice(newSlice, slice.Slice(i, slice.Len()))
		mapValue.SetMapIndex(key, newSlice)
		return
	}
	slice.Index(i).Set(value)
}
