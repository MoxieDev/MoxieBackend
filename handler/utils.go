package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"

	"moxiechat/data"
)

var logger *log.Logger

/* Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a...)
}*/

// Info will log information with "INFO" prefix to logger
func Info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// Danger will log information with "ERROR" prefix to logger
func Danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

// Warning will log information with "WARNING" prefix to logger
func Warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// ReportStatus is a helper function to return a JSON response indicating outcome success/failure
func ReportStatus(w http.ResponseWriter, success bool, err *data.APIError) {
	var res *data.Outcome
	w.Header().Set("Content-Type", "application/json")
	if success {
		res = &data.Outcome{
			Status: success,
		}
	} else {
		res = &data.Outcome{
			Status: success,
			Error:  err,
		}
	}
	response, _ := json.Marshal(res)
	if _, err := w.Write(response); err != nil {
		Danger("Error writing", response)
	}
}

// convenience function to be chained with another HandlerFunc
// that prints to the console which handler was called.
func logConsole(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}
