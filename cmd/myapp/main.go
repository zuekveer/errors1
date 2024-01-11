package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type errNotFound struct {
	message 	string
}

func (e *errNotFound) Error() string {
	return e.message
}

var (
	errNotFoundMessage = "not found"
	values 		= []string{"aaa", "bbb", "1"}
)

func searchAndConvert(expectedValue string) (int, error) {
	for _, v := range values {
		if v == expectedValue {
			convertedValue, err := strconv.Atoi(v)
			if err != nil {
				return  0, err
			}
			return convertedValue, nil
		}
	}
	return 0, fmt.Errorf("searchAndConvert error: %w", &errNotFound{message: errNotFoundMessage})
}

func main() {
	res, err := searchAndConvert("2")
	if err != nil {
		var enf *errNotFound
		if errors.As(err, &enf) { 
			log.Println("value not found")
			return
		}
		log.Fatal(err)
	}
	log.Println(res)
}