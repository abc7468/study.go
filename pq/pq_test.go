package pq

import (
	"fmt"
	"sort"
)

func ExampleXaseInsensitive_sort() {
	apple := CaseInsensitive([]string{"iPhone", "iPad", "MacBook", "AppStore"})
	sort.Sort(apple)
	fmt.Println(apple)
	// Output:
	// [AppStore iPad iPhone MacBook]
}
