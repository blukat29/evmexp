package util

import "strings"

func RemoveHexPrefix(s string) string {
	if strings.HasPrefix(s, "0x") {
		return s[2:]
	} else {
		return s
	}
}
