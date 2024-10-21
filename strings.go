package golodash

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	splitWordReg         = regexp.MustCompile(`([a-z])([A-Z0-9])|([a-zA-Z])([0-9])|([0-9])([a-zA-Z])|([A-Z])([A-Z])([a-z])`)
	splitNumberLetterReg = regexp.MustCompile(`([0-9])([a-zA-Z])`)
)

// Return Part of a string
func Substring[T ~string](from T, offset int, length uint) T {
	rs := []rune(from)
	size := len(rs)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}

	if offset >= size {
		return GetZeroValue[T]()
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	return T(strings.Replace(string(rs[offset:offset+int(length)]), "\x00", "", -1))
}

// RuneLength is an alias to utf8.RuneCountInString which returns the number of runes in string.
func RuneLength(str string) int {
	return utf8.RuneCountInString(str)
}

// Words splits string into an array of its words.
func Words(str string) []string {
	str = splitWordReg.ReplaceAllString(str, `$1$3$5$7 $2$4$6$8$9`)
	str = splitNumberLetterReg.ReplaceAllString(str, "$1 $2")
	var result strings.Builder
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
		} else {
			result.WriteRune(' ')
		}
	}
	return strings.Fields(result.String())
}

// Capitalize converts the first character of string to upper case and the remaining to lower case.
func Capitalize(str string) string {
	return cases.Title(language.English).String(str)
}

// PascalCase converts string to pascal case. eg."hello world" will be "HelloWorld"
func PascalCase(str string) string {
	items := Words(str)
	for i := range items {
		items[i] = Capitalize(items[i])
	}
	return strings.Join(items, "")
}

// CamelCase converts string to camel case. eg."hello world" will be "helloWorld"
func CamelCase(str string) string {
	items := Words(str)
	for i, item := range items {
		item = strings.ToLower(item)
		if i > 0 {
			item = Capitalize(item)
		}
		items[i] = item
	}
	return strings.Join(items, "")
}

// KebabCase converts string to kebab case. eg. "hello world" will be "hello-world"
func KebabCase(str string) string {
	items := Words(str)
	for i := range items {
		items[i] = strings.ToLower(items[i])
	}
	return strings.Join(items, "-")
}

// Ellipsis trims and truncates a string to a specified length and appends an ellipsis if truncated.
func Ellipsis(str string, length int) string {
	str = strings.TrimSpace(str)

	if len(str) > length {
		if len(str) < 3 || length < 3 {
			return "..."
		}
		return strings.TrimSpace(str[0:length-3]) + "..."
	}

	return str
}
