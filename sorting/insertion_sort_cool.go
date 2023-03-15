package main

import "fmt"

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func main() {
	fmt.Println("Insertion Sort Algorithm")

	arr := []int{1, 5, 2, 4, 3}
	fmt.Println(InsertionSort(arr))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(InsertionSort(arr))

	arr = []int{5, 4, 3, 2, 1}
	fmt.Println(InsertionSort(arr))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(InsertionSort(arr))

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(InsertionSort(arr))
}
