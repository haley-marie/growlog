package internal

import (
	"fmt"
	"log"
)

func ReturnMessageAsError(msg string) error {
	errMsg := fmt.Errorf("%s", msg)
	log.Println(errMsg)
	return errMsg
}

func ReturnFormattedError(msg string, err error) error {
	errMsg := fmt.Errorf("%s: %s", msg, err)
	log.Println(errMsg)
	return errMsg
}
