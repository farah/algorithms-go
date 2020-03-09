package ksmallest

import (
	"fmt"
	"math"
)

func kSmallestPair(arr1 []int, arr2 []int, k int) {

	n1 := len(arr1)
	n2 := len(arr2)

	if k > n1*n2 {
		fmt.Println("k pairs don't exist")
	}

	index2 := make([]int, n1)

	for k > 0 {

		min_sum := math.MaxInt32
		min_index := 0

		for i := 0; i < n1; i++ {
			index := index2[i]

			if index2[i] < n2 && arr1[i]+arr2[index2[i]] < min_sum {

				min_index = i
				min_sum = arr1[i] + arr2[index2[i]]
			}
			_ = index
		}

		fmt.Printf("(%d,%d) ", arr1[min_index], arr2[index2[min_index]])

		index2[min_index] = index2[min_index] + 1

		k--
	}
}
