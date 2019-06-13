package cli

func commandSplit(command ...string) (name []string, parameters []string) {
	if 1 > len(command) {
		return []string(nil), []string(nil)
	}

	for i, token := range command {
		if 1 < len(token) && '-' == token[0] {
			parameters = command[i:]
			break
		}

		name = append(name, token)
        }

	return
}
