package validation

import "github.com/badoux/checkmail"

func ValidateEmail(email string) error {
	// Validate email format
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}

	// Validate email host (domain)
	err = checkmail.ValidateHost(email)
	if err != nil {
		return err
	}

	return nil
}
