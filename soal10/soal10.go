package soal10

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

func soal10a() string {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	str := ""
	tally := map[string]int{
		"huruf": 0,
		"vokal": 0,
		"angka": 0,
		"genap": 0,
	}
	for i := 0; i < 50; i++ {
		number := rand.Intn(100) + 1
		char := strings.Split(alphabets, "")[rand.Intn(26)]
		// fmt.Printf("number: %d char: %s\n", number, char)

		tally["huruf"] += 1
		tally["angka"] += 1

		if number%2 == 0 {
			tally["genap"] += 1
		}

		if strings.Contains("aiueo", char) {
			tally["vokal"] += 1
		}

		str += fmt.Sprintf("%s,%d", char, number)
		if i+1 < 50 {
			str += ","
		}
	}

	fmt.Printf("str: %s\n", str)
	fmt.Printf("result: %+v\n", tally)
	return str
}

func soal10b(str string) string {

	arr := strings.Split(str, ",")
	numbers := map[int]bool{}
	chars := map[string]bool{}

	for _, value := range arr {
		if n, err := strconv.Atoi(value); err != nil {
			chars[value] = true
		} else {
			numbers[n] = true
		}
	}

	keys_for_numbers := []int{}
	keys_for_chars := []string{}

	for key, _ := range numbers {
		keys_for_numbers = append(keys_for_numbers, key)
	}

	for key, _ := range chars {
		keys_for_chars = append(keys_for_chars, key)
	}

	slices.Sort(keys_for_chars)
	slices.SortFunc(keys_for_numbers, func(v1, v2 int) int {
		if v1 < v2 {
			return 1
		} else if v1 > v2 {
			return -1
		}
		return 0
	})

	result := strings.Trim(fmt.Sprintf("%v", keys_for_numbers), "[]")
	return fmt.Sprintf("%s,%s", strings.Join(strings.Fields(result), ","), strings.Join(keys_for_chars, ","))
}

func soal10c(str string) string {
	arr := strings.Split(str, ",")
	numbers := []int{}
	chars := []string{}

	for _, value := range arr {
		if n, err := strconv.Atoi(value); err != nil {
			chars = append(chars, value)
		} else {
			numbers = append(numbers, n)
		}
	}

	slices.Sort(chars)
	slices.SortFunc(numbers, func(v1, v2 int) int {
		if v1 < v2 {
			return 1
		} else if v1 > v2 {
			return -1
		}
		return 0
	})

	result := []string{}
	for i, v := range chars {
		result = append(result, fmt.Sprintf("%d%s", numbers[i], v))
	}

	fmt.Println(result)

	fmt.Printf("%+v\n", chars)
	fmt.Printf("%+v\n", numbers)

	return strings.Join(result, ",")
}
