package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 4, 7, 8, 1}
	nums = DeleteDuplicates(nums)

	fmt.Println(nums)

}
