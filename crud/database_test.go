package crud

import (
	"testing"
	"time"
	"strings"
)


func TestAddingNewRecord(t *testing.T){

	db := DB{}
	db.Init()

	all_strings := len(db.data)
	cells := []string{"one", "two", "three"}
	db.create(cells)
	time.Sleep(500 * time.Millisecond)
	if all_strings != len(db.data)-1{
		t.Error("Error in adding")
	}

}


func TestUpdateExistingRecord(t *testing.T){

	db := DB{}
	db.Init()

	cells := []string{"two", "three", "one"}
	db.update(1, cells)
	time.Sleep(500 * time.Millisecond)
	new_record := db.data[1]
	if new_record.Title != cells[0] || new_record.Author != cells[1] || new_record.Genre != cells[2] {
		t.Error("Error in updating")
	}

}

func TestRemoveRecord(t *testing.T){

	db := DB{}
	db.Init()

	all_strings := len(db.data)
	db.delete(1)
	time.Sleep(500 * time.Millisecond)
	if all_strings != len(db.data)+1 {
		t.Error("Error in removing")
	}
}