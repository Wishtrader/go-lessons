package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s, s == nil)

	b := [...]string{"Penn", "Teller", "John"}
	fmt.Println(b, len(b), cap(b))

	a1 := make([]int8, 0) 
	fmt.Println(a1, len(a1), cap(a1), a1 == nil)
	
	a1 = append(a1, []int8{1, 2, 3, 4}...)
	fmt.Println(a1, len(a1), cap(a1))

	a2 := append(a1, 6)
	fmt.Println(a2, len(a2), cap(a2))

	a3 := append(a2, 7)
	fmt.Println(a3, len(a3), cap(a3))
}
