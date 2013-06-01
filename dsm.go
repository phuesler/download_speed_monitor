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
}

func getFile(fileName string) string {
	res, err := http.Get("http://localhost:8080/" + fileName)

	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return string(robots)

}

func serverFiles() {
	http.Handle("/", http.FileServer(http.Dir("./tmp/")))
	go http.ListenAndServe(":8080", nil)
}
