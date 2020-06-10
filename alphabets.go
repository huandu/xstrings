package xstrings

const (
	// ASCIILowercase stores the ASCII lowercase latin alphabet.
	ASCIILowercase = "abcdefghijklmnopqrstuvwxyz"

	// ASCIIUppercase stores the ASCII uppercase latin alphabet.
	ASCIIUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits stores the set of ASCII digits.
	// Do not use this string as a filter for runes,
	// consider using the unicode.IsDigit function instead.
	Digits = "0123456789"

	// HexDigits stores the set of symbols which can represent a hex value.
	HexDigits = "0123456789abcdefABCDEF"

	// OctDigits stores the set of symbols which can represent an octal value.
	OctDigits = "01234567"

	// Punctuation stores the set of ASCII punctuation characters.
	// Do not use this string as a filter for runes,
	// consider using the unicode.IsPunct function instead.
	Punctuation = "#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)
