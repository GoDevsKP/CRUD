package crud

type Book struct {
	Title string
	Author string
	Genre string
}

const BOOK_ENTITY_FIELDS_NUM int = 3