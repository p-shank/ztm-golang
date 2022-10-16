//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type TimeCustom struct {
	hour   int
	minute int
	second int
}

func ParseTime(s string) (TimeCustom, error) {

	if reflect.TypeOf(s).Name() != "string" {
		return TimeCustom{}, errors.New("not a string value")
	}

	timeSlice := strings.Split(s, ":")

	if len(timeSlice) < 3 {
		return TimeCustom{}, errors.New("time string not complete")
	}

	timeStruct := TimeCustom{}

	if hour, err := strconv.ParseInt(timeSlice[0], 10, 32); err == nil {
		timeStruct.hour = int(hour)
	} else {
		return TimeCustom{}, errors.New("hour conversion failed")
	}

	if minute, err := strconv.ParseInt(timeSlice[1], 10, 32); err == nil {
		timeStruct.minute = int(minute)
	} else {
		return TimeCustom{}, errors.New("mintue conversion failed")
	}

	if second, err := strconv.ParseInt(timeSlice[2], 10, 32); err == nil {
		timeStruct.second = int(second)
	} else {
		return TimeCustom{}, errors.New("second conversion failed")
	}

	return timeStruct, nil

}
