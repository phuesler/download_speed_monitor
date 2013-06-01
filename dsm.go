package main

import (
	"crypto/md5"
	"fmt"
	"github.com/kuroneko/gosqlite3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	serverFiles()
	expectedChecksum := "61a48e70c0e1e8f980ade96f394c4301"

	startedAt := time.Now().Unix()

	data := getFile("file")
	h := md5.New()
	io.WriteString(h, string(data))
	actualChecksum := fmt.Sprintf("%x", h.Sum(nil))
	writeFile("tmp/downloaded_file", data)

	finishedAt := time.Now().Unix()

	saveToDb(startedAt, finishedAt, expectedChecksum, actualChecksum, 10, "")

	if expectedChecksum == actualChecksum {
		fmt.Println("OK")
	} else {
		fmt.Println("NOK")
	}

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

func saveToDb(
	startedAt int64, finishedAt int64, md5Source string, md5Target string,
	fileSizeBytes int, errors string) {
	filename := "db/statistics.db"
	db, e := sqlite3.Open(filename)
	if e != nil {
		log.Fatalf("Creating %v failed with error: %v", db, e)
	}

	db.Close()

	if _, e := os.Stat(filename); e != nil {
		log.Fatalf("Checking %v existence failed with error: %v", filename, e)
	}

	//	If new.db already exists and is a valid SQLite3 database this should succeed
	if db, e = sqlite3.Open(filename); e != nil {
		log.Fatalf("Reopening %v failed with error: %v", db, e)
	}
	defer db.Close()
	template := `
  INSERT INTO statistics (
    started_at, finished_at, md5_source, md5_target, file_size_bytes, error_message
  )
  VALUES (
    '%d', '%d', '%s', '%s', '%d', '%s'
  );`
	query := fmt.Sprintf(template,
		startedAt,
		finishedAt,
		md5Source,
		md5Target,
		fileSizeBytes,
		errors)
	if _, e = db.Execute(query); e != nil {
		log.Fatalf("Insert into foo failed with error: %v", e)
	}
}
