package my_database

import (
	"strconv"
	"errors"
	"strings"
	"fmt"
)

//TODO: implement in other way!
func getActions() map[string]string {

	actions := map[string]string{
		"c" : "create",
		"r" : "select",
		"u" : "update",
		"d" : "delete",
	}

	return actions;
}

func ValidateQuery(query []string) bool {

	countOfFields := new(Book).getCountOfFields()

	if (len(query) == 0) {
		return false;
	}
	action := query[0]
	if (action != "create" && action != "update" && action != "delete" && action != "select") {
		return false;
	}

	if (action == "create" && len(query) != (1 + countOfFields)) {
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
		if ((len(query) != 2 + countOfFields) || !isNumber(query[1])) {
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

func HandleQuery(query []string, dataStorage *MyDatabaseWrapper) string {
	//actions := getActions()

	var responseMessage string
	var responseData string
	var err error
	query[0] = strings.ToLower(query[0])
	fmt.Println(query[0])
	switch query[0] {

	case "create":
		err = dataStorage.addOneRecord(query[1:4]);
		responseMessage = "Created"

	case "select":
		if (len(query) == 1) {
			responseData, err = dataStorage.readAllRecords();

		} else {
			id, typeErr := strconv.Atoi(query[1])
			if (typeErr == nil) {
				responseData, err = dataStorage.readOneRecord(id)
			}
		}


	case "update":
		id, typeErr := strconv.Atoi(query[1])
		if (typeErr == nil) {
			err = dataStorage.updateOneRecord(id, query[2:5])

		} else {
			err = errors.New("ID must be integer")
		}

		responseMessage = "Updated"


	case "delete":
		id, typeErr := strconv.Atoi(query[1])
		if (typeErr == nil) {
			err = dataStorage.deleteOneRecord(id)
		}

		responseMessage = "Deleted"
	}
	//err = errors.New("Wrong action")
	if (err == nil) {
		if (query[0] == "select") {
			return responseData
		}

		return responseMessage

	} else {

		return err.Error()

	}
}