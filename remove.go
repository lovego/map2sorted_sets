package map2sorted_sets

import (
	"reflect"
	"sort"
)

func RemoveInt64(m map[int64][]int64, k int64, v int64) {
	slice, ok := m[k]
	if !ok || len(slice) == 0 {
		return
	}
	if i := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= v
	}); i < len(slice) && slice[i] == v {
		m[k] = append(slice[0:i], slice[i+1:]...)
	}
}

func RemoveStructByIntField(m, k interface{}, target int64, field string) {
	mapValue, key := reflect.ValueOf(m), reflect.ValueOf(k)
	slice := mapValue.MapIndex(key)
	if !slice.IsValid() {
		return
	}
	if i := sort.Search(slice.Len(), func(i int) bool {
		return slice.Index(i).FieldByName(field).Int() >= target
	}); i < slice.Len() && slice.Index(i).FieldByName(field).Int() == target {
		mapValue.SetMapIndex(key, reflect.AppendSlice(
			slice.Slice(0, i), slice.Slice(i+1, slice.Len()),
		))
	}
}

func RemoveStructByStringField(m, k interface{}, target string, field string) {
	mapValue, key := reflect.ValueOf(m), reflect.ValueOf(k)
	slice := mapValue.MapIndex(key)
	if !slice.IsValid() {
		return
	}
	if i := sort.Search(slice.Len(), func(i int) bool {
		return slice.Index(i).FieldByName(field).String() >= target
	}); i < slice.Len() && slice.Index(i).FieldByName(field).String() == target {
		mapValue.SetMapIndex(key, reflect.AppendSlice(
			slice.Slice(0, i), slice.Slice(i+1, slice.Len()),
		))
	}
}
