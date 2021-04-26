package main

import (
	"fmt"
)

func MaxSubArr(arr []int) (subarr []int, Msum int) {
	if len(arr) == 0 && len(arr) == 1 {
		return
	}
	Msum = 0
	var low int = 0
	var high int = 0
	var sum = 0
	for i, v := range arr {
		sum += v
		if sum <= 0 {
			sum = 0
			low = i + 1
		}
		if sum > Msum {
			Msum = sum
			high = i
		}
	}
	return arr[low:high], Msum
}

func main() {
	arr := []int{-7, 56, 1, -5, -43}
	maxsum, subarr := MaxSubArr(arr)
	fmt.Println(maxsum)
	fmt.Println(subarr)
}
