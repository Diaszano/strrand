package strrand_test

import (
	"fmt"
	"math/rand/v2"
	"regexp"
	"testing"

	. "github.com/Diaszano/strrand"
)

// TestPredefinedStringGenerators validates various string generator functions provided by the strrand package.
// Each generator is tested against its expected output format, including invalid and typical input lengths.
//
//gocyclo:ignore
func TestPredefinedStringGenerators(t *testing.T) {
	// Define test cases for different string generator functions and their respective regex patterns.
	testCases := map[string]struct {
		f     func(int) string // String generator function
		regex *regexp.Regexp   // Regex to validate the output
	}{
		"Binary":           {f: Binary, regex: regexp.MustCompile(`^(?i)[01]+$`)},
		"Octal":            {f: Octal, regex: regexp.MustCompile(`^(?i)[0-7]+$`)},
		"Decimal":          {f: Decimal, regex: regexp.MustCompile(`^(?i)[0-9]+$`)},
		"Hexadecimal":      {f: Hexadecimal, regex: regexp.MustCompile(`^(?i)[0-9a-f]+$`)},
		"CapitalLetters":   {f: CapitalLetters, regex: regexp.MustCompile(`^(?i)[A-Z]+$`)},
		"LowercaseLetters": {f: LowercaseLetters, regex: regexp.MustCompile(`^(?i)[a-z]+$`)},
		"SpecialLetters":   {f: SpecialLetters, regex: regexp.MustCompile(`^(?i)[!@#$%^&*\(\)\-_=+\[\]{}|;:',.<>?/` + "`~]+$")},
		"Base62":           {f: Base62, regex: regexp.MustCompile(`^(?i)[0-9A-Z]+$`)},
		"Base64":           {f: Base64, regex: regexp.MustCompile(`^(?i)[0-9A-Za-z+/]+$`)},
		"Letters":          {f: Letters, regex: regexp.MustCompile(`^(?i)[A-Za-z]+$`)},
		"DefaultString":    {f: DefaultString, regex: regexp.MustCompile(`^(?i)[0-9A-Za-z!@#$%^&*\(\)\-_=+\[\]{}|;:',.<>?/` + "`~]+$")},
	}

	// Iterate through each test case.
	for name, config := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Subtest: Validate edge cases for invalid input lengths.
			t.Run("InvalidInputs", func(t *testing.T) {
				t.Run("NegativeLength", func(t *testing.T) {
					result := config.f(-1)
					if len(result) != 0 {
						t.Errorf("Negative length: expected empty string, got %q", result)
					}
				})

				t.Run("ZeroLength", func(t *testing.T) {
					result := config.f(0)
					if len(result) != 0 {
						t.Errorf("Zero length: expected empty string, got %q", result)
					}
				})
			})

			// Subtest: Validate typical input lengths.
			t.Run("TypicalLengths", func(t *testing.T) {
				cases := []struct {
					length int
				}{
					{length: 4},
					{length: 16},
					{length: 1024},
					{length: 8192},
					{length: 32768},
				}

				for _, tc := range cases {
					t.Run(fmt.Sprintf("Length-%d", tc.length), func(t *testing.T) {
						t.Parallel()
						result := config.f(tc.length)

						// Verify output length matches the input length.
						if len(result) != tc.length {
							t.Errorf("Expected length %d, got %d", tc.length, len(result))
						}

						// Validate the output format using regex.
						if !config.regex.MatchString(result) {
							t.Errorf("Result contains invalid characters: %q", result)
						}
					})
				}
			})

			// Subtest: Validate outputs for random input lengths.
			for range 30 {
				t.Run("RandomLength", func(t *testing.T) {
					t.Parallel()
					length := rand.IntN(100000)
					result := config.f(length)

					// Verify output length matches the random input length.
					if len(result) != length {
						t.Errorf("Random length %d: expected length %d, got %d", length, length, len(result))
					}

					// Validate the output format using regex.
					if length > 0 && !config.regex.MatchString(result) {
						t.Errorf("Result contains invalid characters: %q", result)
					}
				})
			}
		})
	}
}

// TestString validates the behavior of the String function.
//
//gocyclo:ignore
func TestString(t *testing.T) {
	// Test cases for the default charset.
	t.Run("DefaultCharset", func(t *testing.T) {
		cases := []struct {
			length int
		}{
			{length: -1}, // Invalid length
			{length: 0},  // Zero length
			{length: 4},
			{length: 16},
			{length: 1024},
			{length: 8192},
			{length: 32768},
		}
		regex := regexp.MustCompile(`^(?i)[0-9A-Za-z!@#$%^&*\(\)\-_=+\[\]{}|;:',.<>?/` + "`~]+$")

		for _, tc := range cases {
			t.Run(fmt.Sprintf("Length-%d", tc.length), func(t *testing.T) {
				t.Parallel()

				result := String(tc.length)

				// Verify the output length matches the input length.
				if tc.length <= 0 {
					if len(result) != 0 {
						t.Errorf("Expected empty string for length %d, got %q", tc.length, result)
					}

					return
				}

				if len(result) != tc.length {
					t.Errorf("Expected length %d, got %d", tc.length, len(result))
				}

				// Validate the output using the regex for the default charset.
				if !regex.MatchString(result) {
					t.Errorf("Result contains invalid characters: %q", result)
				}
			})
		}
	})

	// Test cases for custom charsets.
	t.Run("CustomCharset", func(t *testing.T) {
		customCharset := "abc12345"
		regex := regexp.MustCompile(`^[a-c1-5]*$`)
		cases := []struct {
			length int
		}{
			{length: -1}, // Invalid length
			{length: 0},  // Zero length
			{length: 4},
			{length: 16},
			{length: 1024},
			{length: 8192},
			{length: 32768},
		}

		for _, tc := range cases {
			t.Run(fmt.Sprintf("Length-%d", tc.length), func(t *testing.T) {
				t.Parallel()
				result := String(tc.length, customCharset)

				// Verify the output length matches the input length.
				if tc.length <= 0 {
					if len(result) != 0 {
						t.Errorf("Expected empty string for length %d, got %q", tc.length, result)
					}

					return
				}

				if len(result) != tc.length {
					t.Errorf("Expected length %d, got %d", tc.length, len(result))
				}

				// Validate the output using the regex for the custom charset.
				if tc.length > 0 && !regex.MatchString(result) {
					t.Errorf("Result contains invalid characters: %q", result)
				}
			})
		}
	})

	// Test random lengths with the default charset.
	t.Run("RandomLengths", func(t *testing.T) {
		regex := regexp.MustCompile(`^(?i)[0-9A-Za-z!@#$%^&*\(\)\-_=+\[\]{}|;:',.<>?/` + "`~]+$")
		for range 30 {
			t.Run("Random", func(t *testing.T) {
				t.Parallel()

				length := rand.IntN(100000)
				result := String(length)

				// Verify the output length matches the random input length.
				if len(result) != length {
					t.Errorf("Expected length %d, got %d", length, len(result))
				}

				// Validate the output using the regex for the default charset.
				if length > 0 && !regex.MatchString(result) {
					t.Errorf("Result contains invalid characters: %q", result)
				}
			})
		}
	})
}
