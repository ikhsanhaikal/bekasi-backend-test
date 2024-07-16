package soal8

import (
	"fmt"
	"slices"
	"strings"
)

func soal8(input []string, kota string) (saran string, ok bool) {
	if len(input) < 1 || len(kota) < 1 {
		return "", false
	}

	loweredCaseKota := strings.ToLower(kota)
	ok = slices.ContainsFunc(input, func(el string) bool {
		return strings.Compare(strings.ToLower(el), loweredCaseKota) == 0
	})

	if !ok {

		// kumpulan_saran_kota := []string{}
		saran = "Saran Kota = "
		for idx, value := range input {
			loweredCaseValue := strings.ToLower(value)
			// fmt.Printf("value: %+v\n\n", value)
			if len(value) > 0 && loweredCaseKota[0] == loweredCaseValue[0] || loweredCaseKota[len(loweredCaseKota)-1] == loweredCaseValue[len(value)-1] {
				// kumpulan_saran_kota = append(kumpulan_saran_kota, value)
				saran += value
				if idx+1 != len(input) {
					saran += ", "
				}
			}

		}
		fmt.Printf("%+v\n", saran)

	}

	return saran, ok
}
