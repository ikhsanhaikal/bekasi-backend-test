package soal8

import "testing"

func TestSoal8(t *testing.T) {
	test_case := []string{"Bandung", "Cimahi", "Ambon", "Jayapura", "Makasar"}
	t.Run("Cimahi", func(t *testing.T) {
		_, ok := soal8(test_case, "Cimahi")
		expected := true

		if ok != expected {
			t.Fatalf("failed Cimahi input")
		}
	})
	t.Run("Bogor", func(t *testing.T) {
		_, ok := soal8(test_case, "Bogor")
		expected := false

		if ok != expected {
			t.Fatalf("failed Bogor input")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		_, ok := soal8(test_case, "")
		expected := false

		if ok != expected {
			t.Fatalf("failed empty input")
		}

	})
}
