package crud

import (
	"regexp"
)


func isNumber(checked string) bool {
	match, _ := regexp.MatchString("[0-9]+", checked)
	return match
}


func SelectValidator(query []string) (bool) {
	if (len(query) != 1 && len(query) != 2) {
		return false
		}

		if (len(query) == 2 && !isNumber(query[1])) {
			return false
		}
	return true
}

func CreateValidator(query []string) (bool) {
	if (len(query) != (1 + BOOK_ENTITY_FIELDS_NUM)){
		return false
		}
	return true
}

func UpdateValidator(query []string) (bool) {
	if ((len(query) != 2 + BOOK_ENTITY_FIELDS_NUM) || !isNumber(query[1])) {
			return false
		}
	return true
}

func DeleteValidator(query []string) (bool) {
	if ((len(query) != 2) || !isNumber(query[1])) {
			return false
		}
	return true
}