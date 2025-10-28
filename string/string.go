package string

import (
	"unicode"
)

// Filter returns a new string with all runes that satisfy the filter function removed.
func Filter(s string, filter func(rune) bool) string {
	var n int
	chars := []rune(s)
	for _, char := range chars {
		if filter(char) {
			continue
		}
		chars[n] = char
		n++
	}
	return string(chars[:n])
}

// HasEmpty checks if there is any empty string in args.
func HasEmpty(args ...string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return true
		}
	}
	return false
}

// NotEmpty checks if all strings in args are non-empty.
func NotEmpty(args ...string) bool {
	return !HasEmpty(args...)
}

// ToCamelCase converts the first character of the string to lowercase.
func ToCamelCase(s string) string {
	for i, v := range s {
		return string(unicode.ToLower(v)) + s[i+1:]
	}
	return ""
}

// Reverse returns the reversed string of s.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Remove removes all occurrences of rmStrings from strings.
func Remove(strings []string, rmStrings ...string) []string {
	if len(rmStrings) == 0 {
		return strings
	}

	mp := make(map[string]struct{}, len(rmStrings))
	for _, s := range rmStrings {
		mp[s] = struct{}{}
	}

	out := make([]string, 0, len(strings))
	for _, s := range strings {
		if _, found := mp[s]; !found {
			out = append(out, s)
		}
	}
	return out
}

// Union returns the union of two []string.
func Union(a, b []string) []string {
	set := make(map[string]struct{}, len(a)+len(b))
	for _, v := range a {
		set[v] = struct{}{}
	}
	for _, v := range b {
		set[v] = struct{}{}
	}
	out := make([]string, 0, len(set))
	for s := range set {
		out = append(out, s)
	}
	return out
}

// TakeOne returns valid string if not empty or later one.
func TakeOne(valid, or string) string {
	if len(valid) > 0 {
		return valid
	}

	return or
}
