// Package linter provides a static analysis tool for validating log messages
package linter

import "unicode"

var sensitivePatterns = [...]string{"password:", "password=", "api_key=", "token:", "token=", "secret:"}

// CheckMessage validates log string against style, localization and security rules
func CheckMessage(msg string) string {
	if msg == "" {
		return ""
	}

	for i := 0; i < len(sensitivePatterns); i++ {
		if containsFold(msg, sensitivePatterns[i]) {
			return "log message contains potentially sensitive data"
		}
	}

	runes := []rune(msg)
	first := runes[0]

	if (first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z') {
		if !unicode.IsLower(first) {
			return "log message must start with a lowercase letter"
		}
	}

	for _, r := range runes {
		if r == '!' || r == '?' {
			return "log message must not contain special characters or emojis"
		}
		if r != '=' && unicode.IsSymbol(r) {
			return "log message must not contain special characters or emojis"
		}
	}

	for _, r := range runes {
		if r > unicode.MaxASCII {
			return "log message must be in English only"
		}

		isEnglish := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
		isDigit := (r >= '0' && r <= '9')
		isAllowedPunct := (r == ' ' || r == '_' || r == '-' || r == ':' || r == '/' || r == '.' || r == '=')

		if !isEnglish && !isDigit && !isAllowedPunct {
			return "log message must be in English only"
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
