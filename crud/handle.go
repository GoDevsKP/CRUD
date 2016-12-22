package crud

import (
	"strconv"
	"errors"
	"strings"
	"fmt"
)


func Handle(query []string, storage *DB) string {
	var responseMessage string
	var responseData string
	var err error
	query[0] = strings.ToLower(query[0])
	fmt.Println(query[0])
	switch query[0] {
        case "create":
            err = storage.create(query[1:4]);
            responseMessage = "\nSuccesfully created...\n"

        case "select":
            if (len(query) == 1) {
                responseData, err = storage.list();

            } else {
                id, typeErr := strconv.Atoi(query[1])
                if (typeErr == nil) {
                    responseData, err = storage.read(id)
                }
            }


        case "update":
            id, typeErr := strconv.Atoi(query[1])
            if (typeErr == nil) {
                err = storage.update(id, query[2:5])

            } else {
                err = errors.New("ID must to be an integer")
            }

            responseMessage = "\nSuccesfully updated...\n"


        case "delete":
            id, typeErr := strconv.Atoi(query[1])
            if (typeErr == nil) {
                err = storage.delete(id)
            }

            responseMessage = "\nSuccesfully deleted...\n"
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