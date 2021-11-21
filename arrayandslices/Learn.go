package arrayandslices

import "fmt"

func Learn() {
	//array fix length
	var arr [2]string
	arr[0] = "A"
	arr[1] = "B"
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// slice
	var s []int = arr2[1:2]
	fmt.Println(s)
	s2 := []int{2, 3, 4, 5, 6, 7, 8, 9}
	s2[0] = 22
	fmt.Println(s2)
	fmt.Printf("Len=%d cap=%d \n", len(s2), cap(s2))
	x := []int{1, 2}
	y := []int{3, 4}
	z := append(x, y...)
	fmt.Println(z)
}
