package responses

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("NICKNAME ALREADY TAKEN")
	}
	if strings.Contains(err, "email") {
		return errors.New("EMAIL ALREADY TAKEN")
	}
	return errors.New("INCORRECT DETAILS")
}
