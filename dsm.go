package main

import (
	"crypto/md5"
	"fmt"
  "os"
	"io"
	"io/ioutil"
	"log"
	"net/http"
  "github.com/kuroneko/gosqlite3"
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
  doTheDb()
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

func doTheDb(){
  filename := "db/statistics.db"
	db, e := sqlite3.Open(filename)
	if e != nil {
		log.Fatalf("Creating %v failed with error: %v", db, e)
	}
	if _, e = db.Execute( "CREATE TABLE foo (id INTEGER PRIMARY KEY ASC, name VARCHAR(10));" ); e != nil {
		log.Fatalf("Create Table foo failed with error: %v", e)
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
	if _, e = db.Execute( "INSERT INTO foo (id,name) VALUES ('1', 'John');" ); e != nil {
		log.Fatalf("Insert into foo failed with error: %v", e)
	}
}
