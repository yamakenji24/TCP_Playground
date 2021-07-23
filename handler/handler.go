package handler

import "log"

func ErrHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}