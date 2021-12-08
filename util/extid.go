package util

import "strings"

const sep = "-"

func EncodeExtId(net, id string) string {
	return net + sep + id
}

func DecodeExtId(extId string) (string, string, bool) {
	parts := strings.Split(extId, sep)
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}
