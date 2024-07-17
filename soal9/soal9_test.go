package soal9

import (
	"fmt"
	"testing"
)

func TestSoal9(t *testing.T) {
	t.Run("bag A", func(t *testing.T) {
		input := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
		resultA := soal9a(input)
		fmt.Printf("result a: %+v\n", resultA)
	})
	t.Run("bag B", func(t *testing.T) {
		input := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
		resultB := soal9b(input)
		fmt.Printf("result b: %+v\n", resultB)
	})
	t.Run("bag C", func(t *testing.T) {
		input := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
		resultC := soal9c(input, []int{9, 1, 6, 4})
		fmt.Printf("result c: %+v\n", resultC)
	})
	t.Run("bag C Part II", func(t *testing.T) {
		input := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
		resultC := soal9c(input, []int{9, 1, 6, 3})
		fmt.Printf("result c: %+v\n", resultC)
	})
	t.Run("bag C Part III", func(t *testing.T) {
		input := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}
		resultC := soal9c(input, []int{9, 1, 6, 2, 2})
		fmt.Printf("result c: %+v\n", resultC)
	})
	arr := []int{1, 2, 3, 4}

	fmt.Printf("%+v\n", remove(arr, 0))
	fmt.Printf("%+v\n", remove(arr, len(arr)-1))
	fmt.Printf("%+v\n", remove(arr, 2))
	// remove 1th
	// remove nth
	// remove 1 > x < nth

}
