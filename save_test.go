package map2sorted_sets

import "fmt"

func ExampleSaveInt64() {
	var m = make(map[int64][]int64)
	for _, v := range []int64{3, 6, 2, 8, 9, 1, 5, 2, 4, 6, 7} {
		SaveInt64(m, 8, v)
	}
	fmt.Println(m)
	// Output: map[8:[1 2 3 4 5 6 7 8 9]]
}

type testStruct struct {
	Id   int64
	Name string
}

func ExampleSaveStructByIntField() {
	var m = make(map[int64][]testStruct)
	for _, v := range []testStruct{
		{Id: 3, Name: `name3`}, {Id: 6, Name: `name_6`}, {Id: 2, Name: `name_2`},
		{Id: 5, Name: `name5`}, {Id: 4, Name: `name4`}, {Id: 7, Name: `name7`},
		{Id: 2, Name: `name-2`}, {Id: 9, Name: `name9`}, {Id: 1, Name: `name1`},
		{Id: 2, Name: `name2`}, {Id: 8, Name: `name8`}, {Id: 6, Name: `name6`},
	} {
		SaveStructByIntField(m, int64(8), v, `Id`)
	}
	fmt.Println(m)
	// Output:
	// map[8:[{1 name1} {2 name2} {3 name3} {4 name4} {5 name5} {6 name6} {7 name7} {8 name8} {9 name9}]]
}

func ExampleSaveStructByStringField() {
	var m = make(map[int64][]testStruct)
	for _, v := range []testStruct{
		{Id: 3, Name: `name3`}, {Id: 16, Name: `name6`}, {Id: 12, Name: `name2`},
		{Id: 5, Name: `name5`}, {Id: 4, Name: `name4`}, {Id: 7, Name: `name7`},
		{Id: 22, Name: `name2`}, {Id: 9, Name: `name9`}, {Id: 1, Name: `name1`},
		{Id: 2, Name: `name2`}, {Id: 8, Name: `name8`}, {Id: 6, Name: `name6`},
	} {
		SaveStructByStringField(m, int64(8), v, `Name`)
	}
	fmt.Println(m)
	// Output:
	// map[8:[{1 name1} {2 name2} {3 name3} {4 name4} {5 name5} {6 name6} {7 name7} {8 name8} {9 name9}]]
}
