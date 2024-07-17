package soal9

import (
	"slices"
)

func soal9a(input []int) []int {

	slices.Sort(input)

	unique := []int{}
	note := map[int]bool{}
	for _, value := range input {
		if !note[value] {
			unique = append(unique, value)
			note[value] = true
		}
	}

	return unique
}

func soal9b(input []int) map[int]int {
	slices.Sort(input)

	tally := map[int]int{}
	for _, value := range input {
		tally[value] += 1
	}

	return tally
}

func remove(input []int, index int) []int {
	if index >= len(input) || index < 0 {
		return input
	}

	return append(append([]int{}, input[0:index]...), input[index+1:]...)
}

func soal9c(input, target []int) []int {

	result := []int{}
	for k1, v1 := range input {
		for k2, v2 := range target {
			if v1 == v2 {
				target = remove(target, k2)
				break
			} else {
				result = append(result, v1)
			}
		}
		if len(target) == 0 {
			result = append(result, input[k1+1:]...)
			break
		}
	}

	return result
}

func soal9d(input []int, n int) {
	// kurang mengerti dengan maksud soal
	//  [9, 1, 6, 4, 8,6,6,3,8,2,3,3,4,1,8,2]
	//10,10,10,5, 8,6,6,3,8,2,3,3,4,1,8,2
}
