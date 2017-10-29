package sorting

import (
	"sort"
	"fmt"
)

func ExampleSortBasic() {

	integers := []int{4, 3, 8, 9, 0, 2, 1, 3, 4}

	sort.Ints(integers)

	fmt.Printf("%v\n", integers)
	fmt.Println(sort.IntsAreSorted(integers))
	fmt.Printf("9 is at pos %d\n", sort.SearchInts(integers, 9))

	// Output:
	// [0 1 2 3 3 4 4 8 9]
	// true
	// 9 is at pos 8

}
