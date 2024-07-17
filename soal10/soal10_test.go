package soal10

import (
	"fmt"
	"testing"
)

func TestSoal10(t *testing.T) {
	t.Run("bag A", func(t *testing.T) {

		str := soal10a()

		fmt.Println(str)

		// fmt.Printf("%s\n", soal10c(str))

		// fmt.Println(soal10b(str))
	})
	t.Run("bag B", func(t *testing.T) {

		str := soal10a()

		fmt.Println()

		fmt.Println(soal10b(str))
	})
	t.Run("bag C", func(t *testing.T) {

		str := soal10a()

		fmt.Println()

		fmt.Printf("%s\n", soal10c(str))

	})
}
