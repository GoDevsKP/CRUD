package crud

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"bufio"
	"fmt"
)


func LoadFromFile(filename string) map[int]Book {

	var result map[int]Book
	data, err := ioutil.ReadFile(filename)

	if (err == nil) {
		json.Unmarshal(data, &result)
	} else {
		return make(map[int]Book)
	}
	return result
}

func SaveToFile(filename string, data map[int]Book) error {
	fileHandle, _ := os.Create(filename)
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