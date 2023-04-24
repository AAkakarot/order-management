package utility

import "log"

// TODO: to use if enhance error handling
func ErrorHandler(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
