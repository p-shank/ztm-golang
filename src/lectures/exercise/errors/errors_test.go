package timeparse

import (
	"fmt"
	"testing"
)

func TestParseTime(t *testing.T) {
	timeStr := "14:20:34"
	if timeStruct, err := ParseTime(timeStr); err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(timeStruct)
	}
}
