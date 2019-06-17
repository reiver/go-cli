package cli

import (
	"strings"
)

func unserialize(command string) []string {
	result := strings.Split(command, "\u001F")
	if 1 == len(result) && "" == result[0] {
		return []string{}
	}
	return result
}
