// Package linter provides a static analysis tool for validating log messages
package linter

import (
	"unicode"
)

var sensitiveKeywords = [...]string{"password", "api_key", "token", "secret"}

// CheckMessage validates log string against style, localization and security rules
func CheckMessage(msg string) string {
	if msg == "" {
		return ""
	}

	runes := []rune(msg)

	if !unicode.IsLower(runes[0]) {
		return "log message must start with a lowercase letter"
	}

	hasSpecialOrEmoji := false
	msgLen := len(msg)

	for i := 0; i < msgLen; i++ {
		b := msg[i]

		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			continue
		}

		if b >= '0' && b <= '9' {
			continue
		}

		if b == ' ' || b == '_' || b == '-' || b == ':' || b == '/' || b == '.' || b == '=' {
			continue
		}

		r := runes[0]
		for _, curRune := range runes {
			if string(curRune) == string(b) {
				r = curRune
				break
			}
		}

		if unicode.IsSymbol(r) || b == '!' || b == '?' {
			hasSpecialOrEmoji = true
			continue
		}

		return "log message must be in English only"
	}

	if hasSpecialOrEmoji {
		return "log message must not contain special characters or emojis"
	}

	for i := 0; i < len(sensitiveKeywords); i++ {
		if containsFold(msg, sensitiveKeywords[i]) {
			return "log message contains potentially sensitive data"
		}
	}

	return ""
}

func containsFold(s, substr string) bool {
	subLen := len(substr)
	if subLen == 0 {
		return true
	}
	if len(s) < subLen {
		return false
	}

	for i := 0; i <= len(s)-subLen; i++ {
		match := true
		for j := 0; j < subLen; j++ {
			c1 := s[i+j]
			c2 := substr[j]
			if c1 == c2 {
				continue
			}
			if c1 >= 'A' && c1 <= 'Z' {
				c1 += 'a' - 'A'
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 += 'a' - 'A'
			}
			if c1 != c2 {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
