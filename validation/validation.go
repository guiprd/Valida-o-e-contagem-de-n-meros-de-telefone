package validation

import (
	"fmt"
	"regexp"
)

func Validation(phones []string) []bool {
	r := make([]bool, len(phones))
	// zeroFilter, err := regexp.Compile("^0{2}")
	// plusFilter, err := regexp.Compile("^[+]")
	specialCharacterFilter, err := regexp.Compile("[$&,:;=?@#|'<>.^*()%!-]")
	if err != nil {
		fmt.Println(err.Error())
	}
	normalNumberFilter, err := regexp.Compile("[+][1-9]{2}[0-9]{5,10}")
	if err != nil {
		fmt.Println(err.Error())
	}

	for i, number := range phones {
		isSpecialCharacter := specialCharacterFilter.MatchString(number)
		if isSpecialCharacter {
			r[i] = !isSpecialCharacter
			continue
		}
		isNormalNumberFilter := normalNumberFilter.MatchString(number)
		if isNormalNumberFilter {
			r[i] = isNormalNumberFilter
		}
	}
	return r
}
