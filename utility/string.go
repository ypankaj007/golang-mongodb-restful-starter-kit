package utility

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/context"
)

func GetLoggedInUserId(r *http.Request) string {
	id := context.Get(r, "userId")
	return fmt.Sprintf("%v", id)
}

// ValidateRequireAndLengthAndRegex is used to validate any input data but in string type
// @params value is the input value
// @params isRequired definds the wheather the input value required or not
// @params minLength definds minimum length of the input value, 0 value definds no min length check
// @params maxLength definds maximum length of the input value,  0 value definds no max length check
// @params maxLength definds regex of the input value, "" value definds no regex required
// @returns error if any
func ValidateRequireAndLengthAndRegex(value string, isRequired bool, minLength, maxLength int, regex, fieldName string) error {

	length := len(value)
	Re := regexp.MustCompile(regex)
	if isRequired == true && length < 1 {
		return errors.New(fieldName + " is Required")
	}
	if minLength != 0 && length > 1 && length < minLength {
		return errors.New(fieldName + " must be min " + strconv.Itoa(minLength))
	}

	if maxLength != 0 && length > 1 && length > maxLength {
		return errors.New(fieldName + " must be max " + strconv.Itoa(maxLength))
	}

	if !Re.MatchString(value) {
		return errors.New("Invalid " + fieldName)
	}

	return nil
}
