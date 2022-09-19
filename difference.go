package main

func differenceValue(list []int64) int64 {
	length := len(list)
	if length < 2 {
		return 0
	}

	max, min := list[0], list[0]

	for i := 0; i < length; i++ {
		if list[i] < min {
			min = list[i]
		}
		if list[i] > max {
			max = list[i]
		}
	}
	return max - min
}
