package strrand

import (
	ibytes "bytes"
	"crypto/rand"
	"encoding/binary"
)

const (
	// BinaryCharset defines a binary character set (0 and 1).
	BinaryCharset = "01"
	// OctalCharset defines an octal character set (0-7).
	OctalCharset = "01234567"
	// DecimalCharset defines a decimal character set (0-9).
	DecimalCharset = "0123456789"
	// HexadecimalCharset defines a hexadecimal character set (0-9, a-f).
	HexadecimalCharset = "0123456789abcdef"
)

const (
	// UppercaseCharset defines uppercase alphabetic characters (A-Z).
	UppercaseCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// LowercaseCharset defines lowercase alphabetic characters (a-z).
	LowercaseCharset = "abcdefghijklmnopqrstuvwxyz"
	// SpecialCharset defines special characters.
	SpecialCharset = "!@#$%^&*()-_=+[]{}|;:',.<>?/`~"
)

const (
	// AlphabetCharset combines uppercase and lowercase alphabetic characters.
	AlphabetCharset = UppercaseCharset + LowercaseCharset
	// Base62Charset defines a Base62 character set (0-9, A-Z, a-z).
	Base62Charset = DecimalCharset + AlphabetCharset
	// Base64Charset defines a Base64 character set (0-9, A-Z, a-z, +, /).
	Base64Charset = Base62Charset + "+/"
	// DefaultCharset defines a default character set (Base62 + special characters).
	DefaultCharset = Base62Charset + SpecialCharset
)

// Binary generates a random string of the specified length using the binary charset (01).
func Binary(length uint32) string {
	return random(length, BinaryCharset)
}

// Octal generates a random string of the specified length using the octal charset (01234567).
func Octal(length uint32) string {
	return random(length, OctalCharset)
}

// Decimal generates a random string of the specified length using the decimal charset (0123456789).
func Decimal(length uint32) string {
	return random(length, DecimalCharset)
}

// Hexadecimal generates a random string of the specified length using the hexadecimal charset (0123456789abcdef).
func Hexadecimal(length uint32) string {
	return random(length, HexadecimalCharset)
}

// CapitalLetters generates a random string of the specified length using uppercase letters.
func CapitalLetters(length uint32) string {
	return random(length, UppercaseCharset)
}

// LowercaseLetters generates a random string of the specified length using lowercase letters.
func LowercaseLetters(length uint32) string {
	return random(length, LowercaseCharset)
}

// SpecialLetters generates a random string of the specified length using special characters.
func SpecialLetters(length uint32) string {
	return random(length, SpecialCharset)
}

// Base62 generates a random string of the specified length using the base62 charset (0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz).
func Base62(length uint32) string {
	return random(length, Base62Charset)
}

// Base64 generates a random string of the specified length using the base64 charset (0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/).
func Base64(length uint32) string {
	return random(length, Base64Charset)
}

// Letters generates a random string of the specified length using both uppercase and lowercase letters.
func Letters(length uint32) string {
	return random(length, AlphabetCharset)
}

// DefaultString generates a random string of the specified length using the default charset (base62 + special characters).
func DefaultString(length uint32) string {
	return random(length, DefaultCharset)
}

// random generates a random string of the specified length using the provided charset.
func random(length uint32, charset string) string {
	var buffer ibytes.Buffer
	buffer.Grow(int(length))

	charsetRune := []rune(charset)
	charsetRuneLength := uint32(len(charsetRune))

	for range length {
		index := binary.BigEndian.Uint32(bytes(4)) % charsetRuneLength
		buffer.WriteRune(charsetRune[index])
	}

	return buffer.String()
}

// bytes generates a byte slice of the specified length filled with random data.
func bytes(length uint32) []byte {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// String generates a random string of the specified length using a custom charset if provided,
// otherwise, it uses the default charset.
func String(length uint32, customCharset ...string) string {
	if len(customCharset) == 0 {
		return DefaultString(length)
	}

	return random(length, customCharset[0])
}
