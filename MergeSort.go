package main

func MergeSort(arr []int) {
	if len(arr) > 1 {
		q := len(arr) / 2
		Left1 := arr[:q]
		Right1 := arr[q:]

		MergeSort(Left1)
		MergeSort(Right1)
		Merge(arr, Left1, Right1)
	}
}

func Merge(arr []int, Left []int, Right []int) {
	Left2 := make([]int, len(Left))
	Right2 := make([]int, len(Right))

	copy(Left2, Left)
	copy(Right2, Right)

	for i, j, k := 0, 0, 0; i < len(Left2) && j < len(Right2); k++ {
		if Left2[i] <= Right2[j] {
			arr[k] = Left2[i]
			i++
		} else {
			arr[k] = Right2[j]
			j++
		}
	}
}
