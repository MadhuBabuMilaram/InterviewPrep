package main

import "fmt"

// buildPrefixSum calculates the prefix sum array
func buildPrefixSum(arr []int) []int {
	n := len(arr)
	prefix := make([]int, n)
	prefix[0] = arr[0]

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}

	return prefix
}

func main() {
	arr := []int{2, 4, 6, 8}

	fmt.Print("Original array: ")
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	prefix := buildPrefixSum(arr)

	fmt.Print("Prefix sum array: ")
	for _, v := range prefix {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
