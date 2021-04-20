package validation_test

import (
	"testing"
	"workwork/src/validation"
)

func TestIsValidKey_WhenKeyIsValid_ReturnsTrue(t *testing.T) {
	testCases := []string{
		"foo",
		"foo_bar",
		"foo.bar",
		"foo.bar_bar",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			isValid := validation.IsValidKey(testCase)
			if isValid == false {
				t.Fatalf("expected '%s' to be a valid key but it is not valid", testCase)
			}
		})
	}
}
