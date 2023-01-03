package strings

import "strings"

// Join concatenates the elements of a slice into a string using a separator
func Join(slice []string, sep string) string {
	return strings.Join(slice, sep)
}
