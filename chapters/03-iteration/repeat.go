package iteration

import "strings"

// Repeat takes a string and repeats it count times.
func Repeat(s string, count int) string {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}