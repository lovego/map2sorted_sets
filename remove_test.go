package map2sorted_sets

import (
	"fmt"
)

func ExampleRemoveInt64() {
	m := map[int64][]int64{8: {1, 2, 3, 4, 5, 6, 7, 8, 9}}
	for _, v := range []int64{5, 3, 8} {
		RemoveInt64(m, 8, v)
	}
	RemoveInt64(m, 9, 9)
	fmt.Println(m)
	// Output: map[8:[1 2 4 6 7 9]]
}

func ExampleRemoveStructByIntField() {
	var m = map[int64][]testStruct{8: {
		{1, `name1`}, {2, `name2`}, {3, `name3`}, {4, `name4`}, {5, `name5`},
		{6, `name6`}, {7, `name7`}, {8, `name8`}, {9, `name9`},
	}}
	for _, v := range []int64{4, 2, 4, 7, 0} {
		RemoveStructByIntField(m, int64(8), v, `Id`)
	}
	RemoveStructByIntField(m, int64(9), int64(9), `Id`)
	fmt.Println(m)
	// Output:
	// map[8:[{1 name1} {3 name3} {5 name5} {6 name6} {8 name8} {9 name9}]]
}

func ExampleRemoveStructByStringField() {
	var m = map[int64][]testStruct{8: {
		{1, `name1`}, {2, `name2`}, {3, `name3`}, {4, `name4`}, {5, `name5`},
		{6, `name6`}, {7, `name7`}, {8, `name8`}, {9, `name9`},
	}}
	for _, v := range []string{"name4", "name2", "name4", "name7", "name0"} {
		RemoveStructByStringField(m, int64(8), v, `Name`)
	}
	RemoveStructByStringField(m, int64(9), "name9", `Name`)
	fmt.Println(m)
	// Output:
	// map[8:[{1 name1} {3 name3} {5 name5} {6 name6} {8 name8} {9 name9}]]
}
