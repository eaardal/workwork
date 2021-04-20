package validation_test

import (
	"testing"
	"workwork/src/validation"
)

func TestIsValidUrl_WhenUrlIsValid_ReturnsTrue(t *testing.T) {
	testCases := []string{
		"http://test.com",
		"https://test.com",
		"http://www.test.com",
		"https://www.test.com",
		"http://test.com/foo",
		"https://test.com/foo",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			isValid := validation.IsValidUrl(testCase)
			if isValid == false {
				t.Fatalf("expected '%s' to be a valid url but it is not valid", testCase)
			}
		})
	}
}

func TestIsValidUrl_WhenUrlIsInvalid_ReturnsFalse(t *testing.T) {
	testCases := []string{
		"",
		"test.com",
		"http//test.com",
		"https//test.com",
		"foo/bar",
		"www.test.com",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			isValid := validation.IsValidUrl(testCase)
			if isValid == true {
				t.Fatalf("expected '%s' to be invalid, but it is valid", testCase)
			}
		})
	}
}

func TestIsValidKey_WhenKeyIsValid_ReturnsTrue(t *testing.T) {
	testCases := []string{
		"foo",
		"foo_bar",
		"foo.bar",
		"foo.bar_bar",
		"foo_foo_bar_bar",
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

func TestIsValidKey_WhenKeyIsInvalid_ReturnsFalse(t *testing.T) {
	testCases := []string{
		"Foo",
		"fooBar",
		"FOO_BAR",
		"FOO",
		"foo.FOO_bar",
		"FOO.foo_bar",
		"foo.foo-bar",
		"foo-bar",
		"-foo_bar",
		"foo_bar-",
		"_foo_bar",
		"foo_bar_",
		"_foo_bar_",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			isValid := validation.IsValidKey(testCase)
			if isValid == true {
				t.Fatalf("expected '%s' to be invalid, but it is valid", testCase)
			}
		})
	}
}
