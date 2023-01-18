/*
The file implements the functions that provide atomic validations.
*/
package validation

import "errors"

/*
The function validates Wildcard DNN.
It shall contain the string "*".
*/
func ValidateWildCardDnn(pattern string) error {
	if !wildcardDnnPattern.MatchString(pattern) {
		return errors.New("the wildcard DNN pattern does not match")
	}

	return nil
}
