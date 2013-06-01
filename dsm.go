package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	serverFiles()
	fmt.Printf("%s", getFile("checksum.md5"))
	fmt.Printf("%s", getFile("file"))
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
