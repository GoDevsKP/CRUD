package my_database

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"bufio"
	"fmt"
)

type MyDatabaseRepository struct {
	filename string
}

func (databaseRepo *MyDatabaseRepository)load() map[int]Book {

	var result map[int]Book
	data, err := ioutil.ReadFile(databaseRepo.filename)

	if (err == nil) {
		json.Unmarshal(data, &result)
	} else {
		return make(map[int]Book)
	}
	return result
}

func (databaseRepo *MyDatabaseRepository)save(data map[int]Book) error {
	fileHandle, _ := os.Create(databaseRepo.filename)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	b, err := json.Marshal(data)
	fmt.Println(b)

	if (err == nil) {
		fmt.Fprintln(writer, string(b))
	}

	writer.Flush()
	return nil
}