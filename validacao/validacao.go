package validacao

import (
	"fmt"
	"regexp"
)

func Validation(phones []string) []string {
	r := make([]string, len(phones))
	validNumber, err := regexp.Compile("^[+]")
	if err != nil {
		fmt.Println(err.Error())
	}
}
