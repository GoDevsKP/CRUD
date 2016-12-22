package crud

import (
	"regexp"
)


func isNumber(checked string) bool {
	match, _ := regexp.MatchString("[0-9]+", checked)
	return match
}


func Check(query []string) bool {

	if (len(query) == 0) {
		return false;
	}
	action := query[0]
	if (action != "create" && action != "update" && action != "delete" && action != "select") {
		return false;
	}

	if (action == "create" && len(query) != (1 + BOOK_ENTITY_FIELDS_NUM)) {
		return false;
	}

	if (action == "select") {
		if (len(query) != 1 && len(query) != 2) {
			return false
		}

		if (len(query) == 2 && !isNumber(query[1])) {
			return false
		}
	}

	if (action == "update") {
		if ((len(query) != 2 + BOOK_ENTITY_FIELDS_NUM) || !isNumber(query[1])) {
			return false
		}
	}

	if (action == "delete") {
		if ((len(query) != 2) || !isNumber(query[1])) {
			return false
		}
	}

	return true
}

