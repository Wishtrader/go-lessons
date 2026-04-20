package main

import "fmt"

func main() {
	transactions := [3]int{10, 5, -7}
	banks := [2]string{"Alfa-BANK", "Techno-BANK"}
	
	arr := []int{1,2,3,4}
	fmt.Println(len(arr), cap(arr))
	
	slice := make([]int8, 5)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	slice2 := append(slice, 1)
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))

	
	slice3 := append(slice2, 2, 3)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))


	fmt.Println(transactions)
	fmt.Println(banks)
}

