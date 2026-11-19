// Sorting in Go. Run: go run sorting.go
// Sorting modifies the original slice.
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. Standard sorts: ascending order.
	nums := []int{5, 2, 8, 1, 3}
	sort.Ints(nums)
	fmt.Println(nums) // [1 2 3 5 8]

	names := []string{"Charlie", "Alice", "Bob"}
	sort.Strings(names)
	fmt.Println(names) // [Alice Bob Charlie]

	prices := []float64{3.5, 1.2, 2.8}
	sort.Float64s(prices)
	fmt.Println(prices) // [1.2 2.8 3.5]

	// 2. Standard descending sort.
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums) // [8 5 3 2 1]

	// 3. Custom sort: people by age ascending.
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 20}}

	// i and j are indexes. Return true if people[i] belongs before people[j].
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("sort.Slice by age:", people) // [{Bob 20} {Alice 25}]

	// Custom sort rules:
	// - Use > instead of < for descending age.
	// - Use people[i].Name < people[j].Name to sort by name.
	// - For age ties, compare names first inside an if:
	//     if people[i].Age == people[j].Age {
	//         return people[i].Name < people[j].Name
	//     }
	//     return people[i].Age < people[j].Age
	// - Use strict comparisons (< or >), never <= or >=.
	// - sort.SliceStable preserves the relative order of equal items.
	// - sort.Slice takes O(n log n) time for constant-time comparisons.
}
