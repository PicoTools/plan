package utils

import "strings"

// StripHexPrefix strips 0x/0X prefixes on hex string
func StripHexPrefix(s string) string {
	s = strings.ReplaceAll(s, "0x", "")
	s = strings.ReplaceAll(s, "0X", "")
	return s
}
