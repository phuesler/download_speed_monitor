package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	serverFiles()
	expectedChecksum := "873d1e9336c929bb418b812f0794212f"
	data := getFile("file")
	h := md5.New()
	io.WriteString(h, string(data))
	actualChecksum := fmt.Sprintf("%x", h.Sum(nil))
	writeFile("tmp/downloaded_file", data)

	fmt.Printf("%s\n%s\n", expectedChecksum, actualChecksum)
}

func getFile(fileName string) []byte {
	res, err := http.Get("http://localhost:8080/" + fileName)

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return body

}

func serverFiles() {
	http.Handle("/", http.FileServer(http.Dir("./tmp/")))
	go http.ListenAndServe(":8080", nil)
}

func writeFile(path string, data []byte) {
	ioutil.WriteFile(path, data, 0755)
}
