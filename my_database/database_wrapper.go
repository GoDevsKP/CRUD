package my_database

import (
	"sync"
	"fmt"
	"errors"
)

type MyDatabaseWrapper struct {
	file sync.Mutex
	data map[int]Book
	repo MyDatabaseRepository
}

func (databaseWrapper *MyDatabaseWrapper) loadDataFromFile() {
	databaseWrapper.file.Lock()
	databaseWrapper.data = databaseWrapper.repo.load()
	databaseWrapper.file.Unlock()
}


func (databaseWrapper *MyDatabaseWrapper) saveDataToFile() {
	databaseWrapper.file.Lock()
	databaseWrapper.repo.save(databaseWrapper.data)
	databaseWrapper.file.Unlock()
}

func (databaseWrapper *MyDatabaseWrapper) Initialize() {

	databaseWrapper.file = sync.Mutex{}
	databaseWrapper.repo = MyDatabaseRepository{filename:"test.json"}
	databaseWrapper.loadDataFromFile()

}

func (databaseWrapper *MyDatabaseWrapper) addOneRecord(fields []string) (error) {

	max_key := 0
	for k,_ := range databaseWrapper.data {
		if (k > max_key) {
			max_key = k
		}
	}
	databaseWrapper.data[max_key + 1] = Book{Title: fields[0], Author: fields[1], Genre: fields[2]}

	go databaseWrapper.saveDataToFile()

	return nil
}

func (databaseWrapper *MyDatabaseWrapper) updateOneRecord(id int, fields []string) (error) {

	if _, ok := databaseWrapper.data[id]; ok {
		databaseWrapper.data[id] = Book{Title: fields[0], Author: fields[1], Genre: fields[2]}
		go databaseWrapper.saveDataToFile()
		return nil
	} else {
		return errors.New("Record not found")
	}
}

func (databaseWrapper *MyDatabaseWrapper) readOneRecord(id int) (string, error) {
	if val, ok := databaseWrapper.data[id]; ok {
		return fmt.Sprintf("id:%d\nTitle: %s\nAuthor: %s\nGenre: %s",
			id, val.Title, val.Author, val.Genre), nil
	} else {
		return "_", errors.New("Record not found")
	}
}

func (databaseWrapper *MyDatabaseWrapper) readAllRecords() (string, error) {
	result := fmt.Sprintf("%s-------%s-------------%s------------%s", "id", "Title", "Author", "Genre")

	for k, v := range databaseWrapper.data {
		result = result + fmt.Sprintf("\n%2d%8s%17s%18s",
			k, v.Title, v.Author, v.Genre)
	}

	return string(result), nil
}

func (databaseWrapper *MyDatabaseWrapper) deleteOneRecord(id int) error {
	if _, ok := databaseWrapper.data[id]; ok {
		delete(databaseWrapper.data, id)
		go databaseWrapper.saveDataToFile()
		return nil
	} else {
		return errors.New("Record not found")
	}
}


