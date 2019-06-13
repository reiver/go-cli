package cli

import (
	"bytes"
)

func serialize(command ...string) (string, error) {
	var buffer bytes.Buffer

	for i, token := range command {
		if 0 != i {
			buffer.WriteRune('\u001F')
		}

		err := encode(&buffer, token)
		if nil != err {
			return buffer.String(), err
		}
	}

	return buffer.String(), nil
}
