package my_database

import (
	"reflect"
)

type Book struct {
	Title     string
	Author  string
	Genre string
}

func (database *Book) getCountOfFields() int {
	typ := reflect.TypeOf((*Book)(nil)).Elem()
	return typ.NumField()
}