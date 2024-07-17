package soal6

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func soal6(path string) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/%s", path))
	if err != nil {
		log.Fatal("failed on request over network: ", err)
		return
	}
	content, _ := io.ReadAll(resp.Body)
	err = os.WriteFile(fmt.Sprintf("%s.txt", path), content, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
