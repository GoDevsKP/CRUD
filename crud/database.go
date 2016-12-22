package crud

import (
	"sync"
	"fmt"
	"errors"
)

type DB struct {
	file sync.Mutex
	data map[int]Book
	filename string
}

func (db *DB) load() {
	db.file.Lock()
	db.data = LoadFromFile(db.filename)
	db.file.Unlock()
}


func (db *DB) save() {
	db.file.Lock()
	SaveToFile(db.filename, db.data)
	db.file.Unlock()
}

func (db *DB) Init() {

	db.file = sync.Mutex{}
	db.filename = "test.json"
	db.load()

}

func (db *DB) create(fields []string) (error) {

	max_key := 0
	for k,_ := range db.data {
		if (k > max_key) {
			max_key = k
		}
	}
	db.data[max_key + 1] = Book{Title: fields[0], Author: fields[1], Genre: fields[2]}

	go db.save()

	return nil
}

func (db *DB) update(id int, fields []string) (error) {

	if _, ok := db.data[id]; ok {
		db.data[id] = Book{Title: fields[0], Author: fields[1], Genre: fields[2]}
		go db.save()
		return nil
	} else {
		return errors.New("Not found")
	}
}

func (db *DB) read(id int) (string, error) {
	if val, ok := db.data[id]; ok {
		return fmt.Sprintf("\nid:%d | Title: %s | Author: %s | Genre: %s\n",
			id, val.Title, val.Author, val.Genre), nil
	} else {
		return "_", errors.New("Not found")
	}
}

func (db *DB) list() (string, error) {
	result := fmt.Sprintf("\n\n%s|%s|%s|%s", "id", "Title", "Author", "Genre")

	for k, v := range db.data {
		result = result + fmt.Sprintf("\n%d | %s | %s | %s",
			k, v.Title, v.Author, v.Genre)
	}

	result = result + "\n\n"

	return string(result), nil
}

func (db *DB) delete(id int) error {
	if _, ok := db.data[id]; ok {
		delete(db.data, id)
		go db.save()
		return nil
	} else {
		return errors.New("Not found")
	}
}


