package util

import (
	"fmt"
	"testing"
)

func TestUnixTime(t *testing.T) {
}

func TestTimeTz(t *testing.T) {
	dateIso := "2022-11-27T23:50:54+0900"
	dateIso = DatetoDateIso(dateIso)
	fmt.Println(dateIso)
	str := GetIsoDateTime(dateIso)
	fmt.Println(dateIso, str)
}
