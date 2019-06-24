package cli

import (
	"github.com/reiver/go-cli/flag"
)

func commandSplit(command ...string) (name []string, parameters []string) {
	if 1 > len(command) {
		return []string(nil), []string(nil)
	}

	for i, token := range command {
		if 1 < len(token) && cliflag.IsFlag(token) {
			parameters = command[i:]
			break
		}

		name = append(name, token)
        }

	return
}
